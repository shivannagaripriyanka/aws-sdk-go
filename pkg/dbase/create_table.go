package dbase

import (
	"aws-sdk-go/pkg/Utilities"
	"aws-sdk-go/pkg/store"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"log"
)

func (s Store) CreateTable(in store.TableIn) (error, *store.CreateTableOutput) {

	if in.TableName == " " {
		Utilities.BadReq(Utilities.InvalidTableName, store.TableIn{TableName: in.TableName})
		return fmt.Errorf("creating a Table Failed"), &store.CreateTableOutput{}
	}

	out, err := s.DynamoClient.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("year"),
			AttributeType: types.ScalarAttributeTypeN,
		}, {
			AttributeName: aws.String("title"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("year"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("title"),
			KeyType:       types.KeyTypeRange,
		}},
		TableName: aws.String(in.TableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})
	if err != nil {
		log.Printf("Couldn't create Table %v . Reason:\n",
			in.TableName)
		log.Fatal(err)
	}
	resp := &store.CreateTableOutput{
		Error:            nil,
		Status:           "Success",
		TableDescription: out.TableDescription,
	}

	log.Printf("Successfully created Table with name %v \n", in.TableName)
	return nil, resp
}
