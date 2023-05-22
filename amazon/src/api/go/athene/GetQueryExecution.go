package athene

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go/aws/request"
)

func GetQueryExecution(Client *athena.Client, queryID string) (*athena.GetQueryExecutionOutput, error) {

	/*
		        This operation performs a GetQueryExecution with the defined access to Athena.
		    Args:
				Client (obj) - The AWS client object with credentials for Athena.
		        queryID (str) - Query execution Id. Example - "32cfdcdb-aaa9-4e4c-9069-6c0fc76b962c".
		    Returns:
		        GetQueryExecution_response (out) - Returns metadata.
	*/

	// Validates mandatory parameters for null and empty values
	if queryID == "" || Client == nil {
		fmt.Printf("Mandatory parameters should not be null or empty.")
		return &athena.GetQueryExecutionOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Sends the request to GetQueryExecution
	resp, err := Client.GetQueryExecution(context.TODO(), &athena.GetQueryExecutionInput{
		QueryExecutionId: aws.String(queryID),
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
