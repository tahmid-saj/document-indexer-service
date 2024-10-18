# document-indexer-service

Document indexer service to generate inverted index mappings (document term matrices) for documents, such that the inverted index mappings can be utilized by search services for read optimization. Developed with Go / Gin, S3, DynamoDB.

<br/>
<br/>

## Directory structure

The directory structure is as follows:
### Root Directory
- **README.md**: Overview of the project and instructions for use.
- **go.mod**: Go module file specifying project dependencies.
- **go.sum**: Hashes of dependencies for module integrity.
- **main.go**: Entry point of the application.

### Directories

#### bucket/
- Contains code for interacting with S3 buckets where documents are stored.

#### conf/
- Configuration files for the application, including S3 and DynamoDB settings.

#### data/
- Directory for storing local data related to the indexing process.

#### dynamodb/
- Code for handling interactions with DynamoDB, which stores the inverted index.

#### indexer/
- Core logic for indexing documents, creating the inverted index mappings.

#### models/
- Data models defining the structure of documents, terms, and indexes.

#### object/
- Contains code for handling object-level operations, likely related to S3.

#### routes/
- Defines API routes using Gin for exposing the indexing service.

#### utils/
- Utility functions and helpers for general use throughout the codebase.

<br/>
<br/>

## Overview

### Design

The high level workflow of the document indexer can be found below. Similar services can be found <a href="https://whimsical.com/web-microservices-6uqvwWZtcBFsNJB2hepGy1">here</a> and below:

#### Document indexer workflow

<img width="518" alt="image" src="https://github.com/user-attachments/assets/daa66cc1-a116-4097-8624-905bc4dc9590">

#### Similar services

<img width="834" alt="image" src="https://github.com/user-attachments/assets/b54088e7-870c-46dd-9cf6-2e5ec27d9d5c">

### Examples

#### Sample inverted index mappings

<img width="1334" alt="image" src="https://github.com/user-attachments/assets/a6c5e3f1-7296-4913-9dba-d8e1daee2d45">

<img width="421" alt="image" src="https://github.com/user-attachments/assets/8a1a422c-be7c-4bc8-9802-2a3825a7de5a">

#### Sample inverted index mappings from API output:

```
// Input
{
    "tableName": "document-indexer-index-mapping",
    "bucketName": "document-indexer-service-test-documents"
}
```

```
// Output
{
    "Ok": true,
    "Response": {
        "\"Lorem": {
            "documentIDs": [
                "lorem_ipsum_3.json"
            ],
            "documentTermFrequencies": [
                1
            ],
            "documentTermLocations": [
                [
                    117
                ]
            ]
        },
        "\"de": {
            "documentIDs": [
                "lorem_ipsum_3.json"
            ],
            "documentTermFrequencies": [
                2
            ],
            "documentTermLocations": [
                [
                    79,
                    150
                ]
            ]
        },
        "'Content": {
            "documentIDs": [
                "lorem_ipsum_2.json"
            ],
            "documentTermFrequencies": [
                1
            ],
            "documentTermLocations": [
                [
                    44
                ]
            ]
        },
        "'lorem": {
            "documentIDs": [
                "lorem_ipsum_2.json"
            ],
            "documentTermFrequencies": [
                1
            ],
            "documentTermLocations": [
                [
                    75
                ]
            ]
        },
        "(The": {
            "documentIDs": [
                "lorem_ipsum_3.json"
            ],
            "documentTermFrequencies": [
                1
            ],
            "documentTermLocations": [
                [
                    84
                ]
            ]
        },
        "(injected": {
            "documentIDs": [
                "lorem_ipsum_2.json"
            ],
            "documentTermFrequencies": [
                1
            ],
            "documentTermLocations": [
                [
                    99
                ]
            ]
        },
        "1.10.32": {
            "documentIDs": [
                "lorem_ipsum_3.json"
            ],
            "documentTermFrequencies": [
                2
            ],
            "documentTermLocations": [
                [
                    75,
                    146
                ]
            ]
        },
        "1.10.32.": {
            "documentIDs": [
                "lorem_ipsum_3.json"
            ],
            "documentTermFrequencies": [
                1
            ],
            "documentTermLocations": [
                [
                    128
                ]
            ]
        },
        "1.10.33": {
            "documentIDs": [
                "lorem_ipsum_3.json"
            ],
            "documentTermFrequencies": [
                2
            ],
            "documentTermLocations": [
                [
                    77,
                    148
                ]
            ]
        }
    }
}
```
