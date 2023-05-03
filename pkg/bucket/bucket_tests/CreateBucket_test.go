package bucket

import (
	"aws-sdk-go/pkg/store"
	"testing"
)

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
