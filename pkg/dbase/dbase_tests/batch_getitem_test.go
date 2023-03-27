package dbase_tests

import (
	"aws-sdk-go/pkg/dbase"
	"aws-sdk-go/pkg/store"
	"log"
	"testing"
)

var TableName = "aws-sdk-go"

func Test_BatchGetItem(t *testing.T) {
	dynamoClient := dbase.Store{
		DynamoClient: store.Dbcfg(),
	}
	err, get := dynamoClient.BatchGetItem(&store.BatchGetItemInput{TableName: TableName})
	if err != nil {
		return
	}
	log.Println("Batch Get Item response", get)
}
