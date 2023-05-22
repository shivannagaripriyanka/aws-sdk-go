package Location

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/location"
	"github.com/aws/aws-sdk-go/aws/request"
)

func CalculateRouteMatrix(Client *location.Client, deparutre []float64, destination []float64, calculator string) (*location.CalculateRouteMatrixOutput, error) {

	/*
		        This operation Post Content with the defined access to the location objects.
		    Args:
				locationClient (obj) - The AWS client object with credentials for location.
				deparutre (array) -  Example - [-180, -90].
		        destination (array) - Example - [180, 90].
				calculator (string)- Name of Calculator - Example "NameCal"
		    Returns:
		        CalculateRouteMatrix_response (out) - Returns the metadata.
	*/

	// Validates mandatory parameters for null and empty values
	if len(deparutre) == 0 || len(destination) == 0 || calculator == "" || Client == nil {
		fmt.Printf("Mandatory parameters should not be null or empty.")
		return &location.CalculateRouteMatrixOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Sends the request to CalculateRouteMatrix
	resp, err := Client.CalculateRouteMatrix(context.TODO(), &location.CalculateRouteMatrixInput{
		CalculatorName:       aws.String(calculator),
		DeparturePositions:   [][]float64{deparutre},
		DestinationPositions: [][]float64{destination},
	})

	// Returns the exception if it exists
	if err != nil {
		fmt.Printf("Exception occurred while CalculateRouteMatrix:" + err.Error())
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

	fmt.Printf("CalculateRouteMatrix operation completed successfully")
	// Returns the CalculateRouteMatrix response
	return resp, err
}
