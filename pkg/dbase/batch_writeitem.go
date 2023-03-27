package dbase

import (
	"aws-sdk-go/pkg/Utilities"
	"aws-sdk-go/pkg/store"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (s *Store) BatchWriteItem(in *store.BatchWriteItemInput) (error, *store.BatchWriteItemOutput) {

	if in.TableName == " " {
		Utilities.BadReq(Utilities.InvalidTableName, store.TableIn{TableName: in.TableName})
		return fmt.Errorf("creating a Table Failed"), &store.BatchWriteItemOutput{}
	}
	out, err := s.DynamoClient.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			"aws-sdk-go": {
				{
					DeleteRequest: &types.DeleteRequest{
						Key: map[string]types.AttributeValue{
							"id": &types.AttributeValueMemberS{Value: "123"},
						},
					},
				},
				{
					PutRequest: &types.PutRequest{
						Item: map[string]types.AttributeValue{
							"id":    &types.AttributeValueMemberS{Value: "234"},
							"name":  &types.AttributeValueMemberS{Value: "dynamobase"},
							"email": &types.AttributeValueMemberS{Value: "dynobase@dynobase.dev"},
						},
					},
				},
			},
			"TableTwo": {
				{
					PutRequest: &types.PutRequest{
						Item: map[string]types.AttributeValue{
							"id":    &types.AttributeValueMemberS{Value: "456"},
							"name":  &types.AttributeValueMemberS{Value: "dynamobase"},
							"email": &types.AttributeValueMemberS{Value: "dynobase@dynobase.dev"},
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	resp := &store.BatchWriteItemOutput{
		Error:            nil,
		Status:           "Success",
		UnprocessedItems: out.UnprocessedItems,
		ResultMetadata:   out.ResultMetadata,
	}
	return nil, resp
}
