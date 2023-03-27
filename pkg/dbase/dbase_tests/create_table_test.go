package dbase_tests

import (
	"aws-sdk-go/pkg/dbase"
	"aws-sdk-go/pkg/store"
	"testing"
)

//func Config() *dynamodb.Client {
//	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	sessionDbClient := dynamodb.NewFromConfig(cfg)
//	return sessionDbClient
//}

func TestCreateTablePass(t *testing.T) {
	dbclient := dbase.Store{
		DynamoClient: Config(),
	}

	table_name := "aws-sdk-go"
	input := store.TableIn{TableName: table_name}
	err, _ := dbclient.CreateTable(input)
	if err != nil {
		t.Logf("Creating table failed with error %v", err)
		t.Fatal("Fail")
	}
	t.Logf("table got Created %v ", table_name)
	t.Log("Pass")
}
