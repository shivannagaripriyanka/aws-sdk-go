package bucket_tests

import (
	"aws-sdk-go/pkg/bucket"
	"aws-sdk-go/pkg/store"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"testing"
)

func Config() *s3.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))

	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	return s3Client
}

func TestCreateBucketPass(t *testing.T) {
	client := bucket.Store{
		Client: Config(),
	}
	bucket_name := "shivannagari-finaltest"
	in := store.BucketIn{BucketName: bucket_name}
	err, s := client.CreateBucket(in)
	if err != nil {
		t.Logf("Creating bucket failed with error %v", err)
		t.Fatal("Fail")
	}
	t.Logf("Bucket got Created %v at %v %v", bucket_name, s.Location, s)
	t.Log("Pass")

}
