package athene

import (
	"aws-sdk-go/amazon/src/api/go/athene"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"log"
	"testing"
)

func TestGetQueryExecution(t *testing.T) {
	// Load AWS configurations
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	// Returns the exception if it exists
	if err != nil {
		log.Fatal(err)
	}

	// Initializes a session
	athClient := athena.NewFromConfig(cfg)
	queryID := "32cfdcdb-aaa9-4e4c-9069-6c0fc76b962c"

	// Sends the request to GetQueryExecution
	_, athErr := athene.GetQueryExecution(athClient, queryID)

	// Returns the exception if it exists
	if athErr != nil {
		t.Log("Exception occurred while Get query execution : ", athErr)
		t.Fatal("Fail")
	}
	t.Log("GetQueryExecution operation completed successfully")
}
