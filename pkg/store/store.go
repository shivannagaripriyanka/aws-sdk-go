package store

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/middleware"
	"io"
	"log"
)

type BucketIn struct {
	Content    io.ReadSeeker
	BucketName string
}
type TableIn struct {
	TableName string
}

type CreateBucketOutput struct {
	Location *string
	Error    error
	Status   string
	Metadata middleware.Metadata
}
type CreateTableOutput struct {
	Error            error
	Status           string
	TableDescription *types.TableDescription
}

type DeleteBucketOutput struct {
	Error    interface{}
	Status   string
	Metadata middleware.Metadata
}

type BatchGetItemInput struct {
	TableName string
}

type BatchGetItemOutput struct {
	//ConsumedCapacity []types.ConsumedCapacity
	Responses map[string][]map[string]types.AttributeValue
	Error     interface{}
	Status    string
	//UnprocessedKeys  map[string]types.KeysAndAttributes
	ResultMetadata middleware.Metadata
}

type BatchWriteItemOutput struct {
	//ConsumedCapacity      []types.ConsumedCapacity
	//ItemCollectionMetrics map[string][]types.ItemCollectionMetrics
	Error            error
	Status           string
	UnprocessedItems map[string][]types.WriteRequest
	ResultMetadata   middleware.Metadata
}

func (b BatchWriteItemOutput) Errr() string {
	//TODO implement me
	panic("implement me")
}

func (b BatchWriteItemOutput) Err() string {
	//TODO implement me
	panic("implement me")
}

type BatchWriteItemInput struct {
	TableName string
}

type ListObjectsInput struct {
}

type ListObjectsOutput struct {
}

type Store interface {
	CreateBucket(in BucketIn) (error, *CreateBucketOutput)
	DeleteBucket(in BucketIn) (error, *DeleteBucketOutput)
	ListObjects(in BucketIn) (error, *ListObjectsOutput)
	BatchGetItem(in BatchGetItemInput) (error, *BatchGetItemOutput)
	BatchWriteItem(in BatchWriteItemInput) (error, *BatchWriteItemOutput)
	Close() error
}

type Config interface {
	S3cfg() *s3.Client
	Dbcfg() *dynamodb.Client
}

func S3cfg() *s3.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	return s3Client
}

func Dbcfg() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}
	svc := dynamodb.NewFromConfig(cfg)
	return svc

}
