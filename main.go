package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type BucketBasics struct {
	S3Client *s3.Client
}

// DeleteObjects deletes a list of objects from a bucket.
func DeleteObjects(bucketName string, objectKeys []string) error {
	var objectIds []types.ObjectIdentifier
	for _, key := range objectKeys {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}

	//load the aws config  with the given default profile name
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		log.Fatal(err)
	}

	// SharedConfigState: session.SharedConfigEnable,
	// todo:  enable it in session

	// Initialising a s3 client
	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})

	if err != nil {
		log.Printf("Couldn't delete objects from bucket %v. Here's why: %v\n", bucketName, err)
	}
	return err
}

// Create Bucket

func CreateBucket(bucketName string, region string) error {
	// var ws bool
	//  ws =  regexp.MustCompile(`\s`).MatchString(bucketName)
	// fmt.Println("ws",ws)
	// if bucketName== " " || !ws  {
	// return errors.New("bucket name is not valid")
	// }
	//load the aws config  with the given default profile name
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		log.Fatal(err)
	}

	// SharedConfigState: session.SharedConfigEnable,
	// todo:  enable it in session

	// Initialising a s3 client
	client := s3.NewFromConfig(cfg)

	//creating a bucket

	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
	if err != nil {
		log.Printf("Couldn't create bucket %v in Region %v. Here's why:\n",
			bucketName, region)
		log.Fatal(err)
	}

	log.Println("Successfully created bucket", bucketName)
	return nil

}

// ListObjects lists the objects in a bucket.
func ListObjects(bucketName string) ([]types.Object, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		log.Fatal(err)
	}

	// SharedConfigState: session.SharedConfigEnable,
	// todo:  enable it in session

	// Initialising a s3 client
	client := s3.NewFromConfig(cfg)

	result, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	var contents []types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
	} else {
		contents = result.Contents
		log.Println(result.Contents)
	}
	return contents, err
}

// DeleteBucket deletes a bucket. The bucket must be empty or an error is returned.
func DeleteBucket(bucketName string) error {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		log.Fatal(err)
	}

	// SharedConfigState: session.SharedConfigEnable,
	// todo:  enable it in session

	// Initialising a s3 client
	client := s3.NewFromConfig(cfg)
	_, err = client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName)})
	if err != nil {
		log.Printf("Couldn't delete bucket %v. Here's why: %v\n", bucketName, err)
	}
	return err
}

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file")
		os.Exit(1)
	}
	bucketName := os.Getenv("BUCKET_NAME")
	// Region := os.Getenv("AWS_REGION")
	// CreateBucket(bucketName, Region)
	// DeleteBucket(bucketName)
	ListObjects(bucketName)

}
