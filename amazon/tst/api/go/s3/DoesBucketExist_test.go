package bucket_test

import (
	"allApi/src/aws/api/golang/bucket"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"testing"
)

func TestHeadBucket(t *testing.T) {
	// Load AWS configurations 
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	// Create a session
	s3Client := s3.NewFromConfig(cfg)
	bucketName := "test-golang-bucket-12"
	// Client Api Call -- HeadBucket
	out, err := bucket.HeadBucket(s3Client, bucketName)
	if err != nil {
		t.Log("HeadBucket failed with error", err, err.Error())
		t.Fatal("Fail")
	}
	fmt.Printf("Bucket already exist %v at %v", bucketName, out)
	t.Log("Pass")
}
