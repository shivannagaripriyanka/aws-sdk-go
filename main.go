package main

import (
	"aws-sdk-go/pkg/bucket"
	"aws-sdk-go/pkg/dbase"
	"aws-sdk-go/pkg/store"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"os"
)

func main() {
	//BucketName := os.Getenv("BUCKET_NAME")
	// DeleteBucket(bucketName)
	//ListObjects(bucketName)

	fmt.Print("###################", os.Getenv("GOMOD"))

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)

	client := bucket.Store{
		Client: s3Client,
	}
	in := store.BucketIn{BucketName: "shivannagari-priyanka"}
	err, s := client.CreateBucket(in)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(s, "###################")

	// GetBatchItems
	TableName := ""
	svc := dynamodb.NewFromConfig(cfg)
	dynamoClient := dbase.Store{DynamoClient: svc}

	//err, get := dynamoClient.BatchGetItem(&store.BatchGetItemInput{TableName: TableName})
	//if err != nil {
	//	return
	//}
	//log.Println("Batch Get Item response", get)

	// WriteBatchItems
	err, write := dynamoClient.BatchWriteItem(&store.BatchWriteItemInput{TableName: TableName})
	if err != nil {
		return
	}
	log.Println("Batch Get Item response", write)
}
