package athene

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go/aws/request"
)

func StopQueryExecution(Client *athena.Client, queryId string) (*athena.StopQueryExecutionOutput, error) {
	/*
		        This operation performs a StopQueryExecution with the defined access to Athena.
		    Args:
				Client (obj) - The AWS client object with credentials for Athena.
		        queryId (str) - Query execution Ids Example - "32cfdcdb-aaa9-4e4c-9069-6c0fc76b962c".
		    Returns:
		        StopQueryExecution_response (out) - Returns metadata.
	*/

	// Validates mandatory parameters for null and empty values
	if queryId == "" || Client == nil {
		fmt.Printf("Mandatory parameters should not be null or empty.")
		return &athena.StopQueryExecutionOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Sends the request to StopQueryExecution
	resp, err := Client.StopQueryExecution(context.TODO(), &athena.StopQueryExecutionInput{
		QueryExecutionId: aws.String(queryId),
	})

	// Returns the exception if it exists
	if err != nil {
		fmt.Printf("Exception occurred while StopQueryExecution:" + err.Error())
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

	fmt.Printf("StopQueryExecution operation completed successfully")
	// Returns the StopQueryExecution response
	return resp, err
}
