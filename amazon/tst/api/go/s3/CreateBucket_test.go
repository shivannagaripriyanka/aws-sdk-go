package s3tests

import (
	"aws-sdk-go/pkg/bucket"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"testing"
)

var s3Client *s3.Client

func init() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	s3Client = s3.NewFromConfig(cfg)
}

func TestCreateBucket(t *testing.T) {
	//
	//cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_ = s3.NewFromConfig(cfg)
	bucketName := "shivannagari-priyanka"
	out, err := bucket.CreateBucket(s3Client, bucketName)
	if err != nil {
		t.Log("Creating bucket failed with error", err, err.Error())
		t.Fatal("Fail")
	}
	fmt.Printf("Bucket got Created %v at %v %v", bucketName, out.Location, out)
	t.Log("Pass")
}

func TestCreateBucketFail(t *testing.T) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	_ = s3.NewFromConfig(cfg)
	bucketName := "shivannagari-priyanka"
	out, err := bucket.CreateBucket(nil, bucketName)
	if err == nil {
		t.Log("Expected error but got ", out)
		t.Fatal("Fail")
	}
	fmt.Printf("Expected error and got error  %v ", err)
	t.Log("Pass")
}
func TestCreateBucketEmptyBname(t *testing.T) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	_ = s3.NewFromConfig(cfg)
	//bucketName := "shivannagari-priyanka"
	out, err := bucket.CreateBucket(nil, "")
	if err == nil {
		t.Log("Expected error but got  ", out)
		t.Fatal("Fail")
	}
	t.Logf("Expected error and got error %v", err)
	t.Log("Pass")
}
