package kendra

import (
	"github.com/aws/aws-sdk-go/service/kendra"
)

func CreateDataSource(Client *kendra.Kendra, input kendra.CreateDataSourceInput) (*kendra.CreateDataSourceOutput, error) {

	result, err := Client.CreateDataSource(input)

}
