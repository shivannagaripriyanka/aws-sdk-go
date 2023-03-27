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

// DeleteBucket deletes a bucket. The bucket must be empty or an error is returned.

func (s Store) DeleteBucket(in store.BucketIn) (error, *store.DeleteBucketOutput) {

	if in.BucketName == " " {
		Utilities.BadRequest(Utilities.InvalidBucketName, store.BucketIn{BucketName: in.BucketName})
		return fmt.Errorf("deleting a Bucket Failed"), &store.DeleteBucketOutput{}
	}

	out, err := s.Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(in.BucketName),
	})

	if err != nil {
		log.Printf("Couldn't delete bucket %v . Reason:\n",
			in.BucketName)
		log.Fatal(err)
	}

	resp := &store.DeleteBucketOutput{
		Error:    nil,
		Status:   "Success",
		Metadata: out.ResultMetadata,
	}

	log.Printf("Successfully deleted bucket with name %v  \n", in.BucketName)
	return nil, resp
}
