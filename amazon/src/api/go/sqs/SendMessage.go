package sqs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws/request"
	"reflect"
)

func SendMessage(Client *sqs.Client, queurl string) (*sqs.SendMessageOutput, error) {
	//Validation check
	var tmp *sqs.Client
	if queurl == "" || Client == nil || reflect.TypeOf(tmp) != reflect.TypeOf(tmp) {
		return &sqs.SendMessageOutput{}, fmt.Errorf(request.InvalidParameterErrCode)
	}

	// Api call
	resp, err := Client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    aws.String(queurl),
		MessageBody: aws.String("hi"),
	})

	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	marshalled, marErr := json.Marshal(resp)
	if marErr != nil {
		fmt.Println(marErr)
		return resp, marErr
	}

	json.Unmarshal(marshalled, &sqs.SendMessageOutput{})
	output := string(marshalled)
	fmt.Println(output)
	return resp, err
}
