package bucket

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/sirupsen/logrus"
)

func CreateBucket(Client *s3.Client, Bucketname string) (*s3.CreateBucketOutput, error) {
	// Validates mandatory parameters for null and empty values
	if Bucketname == "" || Client == nil {
		logrus.WithFields(logrus.Fields{
			"[BucketName]": Bucketname,
			"MethodName":   "Create Bucket",
			"Error":        "Invalid or EmptyBucketName or Client",
		}).Info("End")
		return &s3.CreateBucketOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Sends the request to CreateBucket
	out, err := Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(Bucketname),
	})

	if err != nil {
		fmt.Printf("Exception occurred while Creating the CreateBucket execution :" + err.Error())
		return out, err
	}
	fmt.Printf("CreateBucket execution operation completed successfully")
	// Returns the CreateBucket response
	return out, nil
}
