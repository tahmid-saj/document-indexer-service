package dynamodb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type InvertedIndex struct {
	Term          string
	DocumentTermMatrix DocumentTermMatrix
}

type DocumentTermMatrix struct {
	DocumentIDs []string
	DocumentTermFrequencies []int
	DocumentTermLocations [][]int
}

func ListTables() ([]string, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// create the input configuration instance
	input := &dynamodb.ListTablesInput{}

	var tableNames []string
	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				log.Print(err)
				return nil, err
			}
			log.Print(err)
			return nil, err
		}

		for _, tableName := range result.TableNames {
			tableNames = append(tableNames, *tableName)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}

	return tableNames, nil
}

func CreateTable(tableName string) (*dynamodb.CreateTableOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create the table input with "Term" as the primary key (HASH)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Term"), // Define Term attribute
				AttributeType: aws.String("S"),      // Term is a string (S)
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Term"), // Primary key (HASH)
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	// Create the table
	result, err := svc.CreateTable(input)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}

func AddItem(item InvertedIndex, tableName string) (*dynamodb.PutItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	attributeValue, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(tableName),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}

func AddItemsFromJSON(items []interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	var result *dynamodb.PutItemOutput
	for _, item := range items {
    attributeValue, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
			log.Print(err)
			return nil, err
    }

    // Create item in table Movies
    input := &dynamodb.PutItemInput{
			Item:      attributeValue,
			TableName: aws.String(tableName),
    }

    result, err = svc.PutItem(input)
    if err != nil {
			log.Print(err)
			return nil, err
    }
	}

	return result, nil
}

// Get table items from JSON file
func getItems(fileName string) interface{} {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Print(err)
		return err
	}

	var items interface{}
	json.Unmarshal(raw, &items)
	return items
}

func ReadItem(term, tableName string) (*InvertedIndex, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
    TableName: aws.String(tableName),
    Key: map[string]*dynamodb.AttributeValue{
			"Term": {
				S: aws.String(term),
			},
    },
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if result.Item == nil {
    msg := "Could not find '" + term + "'"
    return nil, errors.New(msg)
	}
			
	var item *InvertedIndex

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return item, nil
}

func UpdateItem(updatedValue InvertedIndex, tableName string) (*dynamodb.UpdateItemOutput, error) {
	// Initialize a session for AWS credentials and config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Prepare DocumentIDs, DocumentTermFrequencies, and DocumentTermLocations for DynamoDB
	documentIDs := []*dynamodb.AttributeValue{}
	for _, docID := range updatedValue.DocumentTermMatrix.DocumentIDs {
		documentIDs = append(documentIDs, &dynamodb.AttributeValue{
			S: aws.String(docID),
		})
	}

	documentTermFrequencies := []*dynamodb.AttributeValue{}
	for _, freq := range updatedValue.DocumentTermMatrix.DocumentTermFrequencies {
		documentTermFrequencies = append(documentTermFrequencies, &dynamodb.AttributeValue{
			N: aws.String(fmt.Sprintf("%d", freq)),
		})
	}

	documentTermLocations := []*dynamodb.AttributeValue{}
	for _, locations := range updatedValue.DocumentTermMatrix.DocumentTermLocations {
		locs := []*dynamodb.AttributeValue{}
		for _, loc := range locations {
			locs = append(locs, &dynamodb.AttributeValue{
				N: aws.String(fmt.Sprintf("%d", loc)),
			})
		}
		documentTermLocations = append(documentTermLocations, &dynamodb.AttributeValue{
			L: locs,
		})
	}

	// Prepare the update expression and the attribute values
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Term": {
				S: aws.String(updatedValue.Term),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#DTM": aws.String("DocumentTermMatrix"),
			"#IDs": aws.String("DocumentIDs"),
			"#Frequencies": aws.String("DocumentTermFrequencies"),
			"#Locations": aws.String("DocumentTermLocations"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ids": {
				L: documentIDs, // list of document IDs
			},
			":freqs": {
				L: documentTermFrequencies, // list of document term frequencies
			},
			":locs": {
				L: documentTermLocations, // list of document term locations
			},
		},
		UpdateExpression: aws.String("SET #DTM.#IDs = :ids, #DTM.#Frequencies = :freqs, #DTM.#Locations = :locs"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	// Execute the UpdateItem request
	result, err := svc.UpdateItem(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteItem(term, tableName string) (*dynamodb.DeleteItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
    Key: map[string]*dynamodb.AttributeValue{
			"Term": {
				S: aws.String(term),
			},
    },
    TableName: aws.String(tableName),
	}

	result, err := svc.DeleteItem(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}