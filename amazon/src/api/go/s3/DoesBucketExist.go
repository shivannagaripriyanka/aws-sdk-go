package bucket

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws/request"
)

func HeadBucket(Client *s3.Client, Bucketname string) (*s3.HeadBucketOutput, error) {

	/*
        This operation is to determine if a bucket exists and you have permission to access it..
    Args:
		Client (obj) - The S3 client object with credentials.
        Bucketname (str) - name of the bucket to find. Example - "Bucketname".
    Returns:
        head_bucket_response (out) - Returns HTTP 200 response with an empty HTTP body.
    
	*/

	// Validates mandatory parameters for null and empty values
	if Bucketname == "" || Client == nil {
		fmt.Printf("Mandatory parameters should not be null or empty.")		
		return &s3.HeadBucketOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Sends the request to HeaddBucket
	resp, err := Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(Bucketname),
	})

	// Returns the exception if it exists
	if err != nil {
		fmt.Printf("Exception occurred while finding the bucket :" + err.Error())
	  	return resp , err
	}

	fmt.Printf("HeadBucket operation completed successfully")
	// Returns the CreateTopic response
	return resp, err
}
