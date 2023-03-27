package bucket_tests

import (
	"aws-sdk-go/pkg/bucket"
	"aws-sdk-go/pkg/store"
	"testing"
)

func TestDeleteBucketPass(t *testing.T) {
	client := bucket.Store{
		Client: Config(),
	}
	bucketName := "shivannagari-final"
	in := store.BucketIn{BucketName: bucketName}
	err, s := client.DeleteBucket(in)
	if err != nil {
		t.Logf("Creating bucket failed with error %v", err)
		t.Fatal("Fail")
	}
	t.Logf("Bucket got Deleted %v at %v %v", s.Metadata, bucketName)
	t.Log("Pass")

}
