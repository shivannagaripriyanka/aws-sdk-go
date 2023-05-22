package bucket

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws/request"
)

func CreateBucket(S3Client *s3.Client, Bucketname string) (*s3.CreateBucketOutput, error) {
	/*
		        This operation creates a new bucket with the defined access to the bucket and objects.
		    Args:
				S3Client (obj) - The AWS client object with credentials for S3.
		        Bucketname (str) - Name of the bucket to be created. Example - "bucketname".
		    Returns:
		        create_bucket_response (out) - Returns the location of the bucket and metadata.

	*/

	// Validates mandatory parameters for null and empty values
	if Bucketname == "" || S3Client == nil {
		fmt.Printf("Mandatory parameters should not be null or empty.")
		return &s3.CreateBucketOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Sends the request to CreateBucket
	out, exception := S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(Bucketname),
	})

	// Returns the exception if it exists
	if exception != nil {
		fmt.Printf("Exception occurred while Creating the Bucket :" + exception.Error())
		return out, exception
	}

	fmt.Printf("Create Bucket operation completed successfully")
	// Returns the CreateBucket response
	return out, nil
}
