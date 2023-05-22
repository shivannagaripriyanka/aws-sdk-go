package athene

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go/aws/request"
)

func StartQueryExecution(Client *athena.Client, query string) (*athena.StartQueryExecutionOutput, error) {

	/*
		        This operation performs a StartQueryExecution with the defined access to Athena.
		    Args:
				Client (obj) - The AWS client object with credentials for Athena.
		        query (str) - Query string. Example - " ".
		    Returns:
		        StartQueryExecution_response (out) - Returns metadata.
	*/

	// Validates mandatory parameters for null and empty values
	if query == "" || Client == nil {
		fmt.Printf("Mandatory parameters should not be null or empty.")
		return &athena.StartQueryExecutionOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Sends the request to GetQueryExecution
	resp, err := Client.StartQueryExecution(context.TODO(), &athena.StartQueryExecutionInput{
		QueryString: aws.String(query),
	})

	// Returns the exception if it exists
	if err != nil {
		fmt.Printf("Exception occurred while GetQueryExecution:" + err.Error())
		return resp, err
	}

	// Makes response in readable format
	marshalled, marErr := json.Marshal(resp)
	if marErr != nil {
		fmt.Println(marErr)
		return resp, marErr
	}

	output := string(marshalled)
	fmt.Println(output)

	fmt.Printf("GetQueryExecution operation completed successfully")
	// Returns the GetQueryExecution response
	return resp, err
}
