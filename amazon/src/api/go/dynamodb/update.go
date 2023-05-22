package dynamoDb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

type Movie struct {
	Id   string
	Name string
}

func (movie Movie) GetKey() map[string]types.AttributeValue {
	title, err := attributevalue.Marshal(movie.Id)
	if err != nil {
		panic(err)
	}
	year, err := attributevalue.Marshal(movie.Name)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": title, "name": year}
}

func UpdateMovie(client *dynamodb.Client, movie Movie, new Movie, Tablename string) (map[string]map[string]interface{}, error) {
	var err error
	var response *dynamodb.UpdateItemOutput
	var attributeMap map[string]map[string]interface{}
	update := expression.Set(expression.Name(movie.Name), expression.Value(new.Name))
	update.Set(expression.Name(movie.Id), expression.Value(movie.Id))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	} else {
		response, err = client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName:                 aws.String(Tablename),
			Key:                       movie.GetKey(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			log.Printf("Couldn't update movie %v. Here's why: %v\n", movie.Id, err)
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &attributeMap)
			if err != nil {
				log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
			}
		}
	}
	return attributeMap, err
}
