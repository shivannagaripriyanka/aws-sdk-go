package kendra

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kendra"
)

func main() {
	// Create a new session using the default region and credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	// Create a Kendra client
	kendraClient := kendra.New(sess)

	// Set up the request parameters for creating an index
	input := &kendra.CreateIndexInput{
		Name:        aws.String("example-index"),
		RoleArn:     aws.String("arn:aws:iam::123456789012:role/KendraRole"),
		Description: aws.String("Example index created with the Kendra API"),
		Edition:     aws.String(kendra.IndexEditionDeveloperEdition),
	}

	// Send the request to create the index
	result, err := kendraClient.CreateIndex(input)

	if err != nil {
		fmt.Println("Error creating index:", err)
		return
	}

	fmt.Println("Index created:", result)
}
