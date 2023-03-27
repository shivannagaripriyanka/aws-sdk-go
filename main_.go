package main

import (
	"aws-sdk-go/pkg/bucket"
	"aws-sdk-go/pkg/dbase"
	"aws-sdk-go/pkg/store"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"

	"context"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)

	client := bucket.Store{
		Client: s3Client,
	}
	in := store.BucketIn{BucketName: "mytestetedbucketname"}
	var createbucketErr, _ = client.CreateBucket(in)
	if createbucketErr != nil {
		fmt.Print(createbucketErr)
		return
	}

	deleteBucketErr, _ := client.DeleteBucket(in)
	if deleteBucketErr != nil {
		fmt.Print(deleteBucketErr)
		return
	}

	sessionDbClient := dynamodb.NewFromConfig(cfg)
	dbclient := dbase.Store{
		DynamoClient: sessionDbClient,
	}
	input := store.TableIn{TableName: "aws-sdk-go"}
	var createdynamoErr, _ = dbclient.CreateTable(input)
	if createdynamoErr != nil {
		fmt.Print(createdynamoErr)
		return
	}
}
