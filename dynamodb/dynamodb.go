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

type TrieNode struct {
	Prefix          string
	FrequentQueries map[string]int
	ChildNodes      []string
	LeafNode        bool
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

	// Create the table input with "Prefix" as the primary key (HASH)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Prefix"), // Define Prefix attribute
				AttributeType: aws.String("S"),      // Prefix is a string (S)
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Prefix"), // Primary key (HASH)
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

func AddItem(item TrieNode, tableName string) (*dynamodb.PutItemOutput, error) {
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

func ReadItem(prefix, tableName string) (*TrieNode, error) {
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
			"Prefix": {
				S: aws.String(prefix),
			},
    },
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if result.Item == nil {
    msg := "Could not find '" + prefix + "'"
    return nil, errors.New(msg)
	}
			
	var item *TrieNode

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return item, nil
}

func UpdateItem(updatedValue TrieNode, tableName string) (*dynamodb.UpdateItemOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Prepare FrequentQueries and ChildNodes for DynamoDB
	frequentQueries := map[string]*dynamodb.AttributeValue{}
	for k, v := range updatedValue.FrequentQueries {
		frequentQueries[k] = &dynamodb.AttributeValue{
			N: aws.String(fmt.Sprintf("%d", v)),
		}
	}

	childNodes := []*dynamodb.AttributeValue{}
	for _, v := range updatedValue.ChildNodes {
		childNodes = append(childNodes, &dynamodb.AttributeValue{
			S: aws.String(v),
		})
	}

	// Prepare the update expression and the attribute values
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Prefix": {
				S: aws.String(updatedValue.Prefix),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#FQ": aws.String("FrequentQueries"),
			"#CN": aws.String("ChildNodes"),
			"#LN": aws.String("LeafNode"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":fq": {
				M: frequentQueries, // map of frequent queries
			},
			":cn": {
				L: childNodes, // list of child nodes
			},
			":ln": {
				BOOL: aws.Bool(updatedValue.LeafNode), // LeafNode value
			},
		},
		UpdateExpression: aws.String("SET #FQ = :fq, #CN = :cn, #LN = :ln"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	// Execute the update
	result, err := svc.UpdateItem(input)
	if err != nil {
		log.Printf("Failed to update item: %v", err)
		return nil, err
	}

	return result, nil
}

func DeleteItem(prefix, tableName string) (*dynamodb.DeleteItemOutput, error) {
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
			"Prefix": {
				S: aws.String(prefix),
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