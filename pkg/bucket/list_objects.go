package bucket

import (
	"aws-sdk-go/pkg/Utilities"
	"aws-sdk-go/pkg/store"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func (s *Store) ListObjects(in store.BucketIn) (error, *store.ListObjectsOutput) {

	if in.BucketName == " " {
		Utilities.BadRequest(Utilities.InvalidBucketName, store.BucketIn{BucketName: in.BucketName})
		return fmt.Errorf("listing Objects Failed"), &store.ListObjectsOutput{}
	}

	//ListObjects in a bucket
	out, err := s.Client.ListObjects(context.TODO(), &s3.ListObjectsInput{
		Bucket: aws.String(in.BucketName),
	})

	if err != nil {
		Utilities.LogInfoWithErrorf("END", in.BucketName, "ListObjects", "Couldn't list objects in bucket")
		//log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", in.BucketName, err)
		log.Fatal(err)

	}
	resp := &store.ListObjectsOutput{
		Error:          nil,
		Status:         "Success",
		Contents:       out.Contents,
		MaxKeys:        0,
		Name:           out.Name,
		ResultMetadata: out.ResultMetadata,
	}
	return nil, resp
}

func abortwithError() {

}
