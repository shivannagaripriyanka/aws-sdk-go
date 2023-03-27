package bucket

import (
	"aws-sdk-go/pkg/Utilities"
	"aws-sdk-go/pkg/store"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"strings"
)

type Store struct {
	Client *s3.Client
}

// Create Bucket
func (s Store) CreateBucket(in store.BucketIn) (error, *store.CreateBucketOutput) {
	//m := make(map[string]bool)
	var ok bool
	a := strings.Trim(in.BucketName, "")
	fmt.Println(a, ok, in.BucketName)
	if in.BucketName == " " {
		Utilities.BadRequest(Utilities.InvalidBucketName, store.BucketIn{BucketName: in.BucketName})
		return fmt.Errorf("creating a BUcket Failed"), &store.CreateBucketOutput{}
	}

	//creating a bucket
	out, err := s.Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(in.BucketName),
		//CreateBucketConfiguration: &types.CreateBucketConfiguration{
		//LocationConstraint: "us-east-2",
		//},
	})
	if err != nil {
		log.Printf("Couldn't create bucket %v . Reason:\n",
			in.BucketName)
		log.Fatal(err)
	}

	resp := &store.CreateBucketOutput{
		Location: out.Location,
		Error:    nil,
		Status:   "Success",
		Metadata: out.ResultMetadata,
	}

	log.Printf("Successfully created bucket with %v at %v \n", in.BucketName, out.Location)
	return nil, resp
}
