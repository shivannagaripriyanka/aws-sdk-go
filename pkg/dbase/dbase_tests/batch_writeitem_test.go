package dbase_tests

import (
	"aws-sdk-go/pkg/dbase"
	"aws-sdk-go/pkg/store"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"testing"
)

func Config() *dynamodb.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc
}
func Test_BatchWriteItem(t *testing.T) {
	dynamoClient := dbase.Store{DynamoClient: Config()}
	err, write := dynamoClient.BatchWriteItem(&store.BatchWriteItemInput{TableName: TableName})
	if err != nil {
		return
	}
	log.Println("Batch Get Item response", write)
}
