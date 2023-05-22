package dynamodb

import (
	dynamoDb "aws-sdk-go/amazon/src/api/go/dynamodb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"testing"
)

func TestUpdateItem(t *testing.T) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}
	dynamoDbClient := dynamodb.NewFromConfig(cfg)
	table_name := "three"

	old := dynamoDb.Movie{
		Id:   "123",
		Name: "abc",
	}
	newval := dynamoDb.Movie{
		Id:   "456",
		Name: "def",
	}
	out, err := dynamoDb.UpdateMovie(dynamoDbClient, old, newval, table_name)

	marshalled, marErr := json.Marshal(out)
	if marErr != nil {
		fmt.Println(marErr)
	}

	json.Unmarshal(marshalled, &dynamodb.UpdateItemOutput{})
	output := string(marshalled)

	if err != nil {
		t.Logf("Got error calling UpdateItem: %s", err)
	}
	fmt.Printf("UpdateItem response %v ", output)

	t.Logf("Item got Updated %v ", table_name)
	t.Log("Pass")
}
