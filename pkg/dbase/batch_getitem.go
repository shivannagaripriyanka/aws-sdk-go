package dbase

import (
	"aws-sdk-go/pkg/Utilities"
	"aws-sdk-go/pkg/store"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Store struct {
	DynamoClient *dynamodb.Client
}

func (s *Store) BatchGetItem(in *store.BatchGetItemInput) (error, *store.BatchGetItemOutput) {

	if in.TableName == " " {
		Utilities.BadReq(Utilities.InvalidTableName, store.TableIn{TableName: in.TableName})
		return fmt.Errorf("creating a Table Failed"), &store.BatchGetItemOutput{}
	}

	out, err := s.DynamoClient.BatchGetItem(context.TODO(), &dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			"aws-sdk-go": {
				Keys: []map[string]types.AttributeValue{
					{
						"id": &types.AttributeValueMemberS{Value: "123"},
					},
					//{
					//	"id": &types.AttributeValueMemberS{Value: "123"},
					//},
				},
			},
			"other-table": {
				Keys: []map[string]types.AttributeValue{
					{
						"id": &types.AttributeValueMemberS{Value: "abc"},
					},
					//{
					//	"id": &types.AttributeValueMemberS{Value: "abd"},
					//},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	resp := &store.BatchGetItemOutput{
		Responses:      out.Responses,
		Error:          nil,
		Status:         "Success",
		ResultMetadata: out.ResultMetadata,
	}
	return nil, resp

}
