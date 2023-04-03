package transcribeservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

func CheckProgress(jobName string) (string, error) {
	svc := transcribeservice.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	input := &transcribeservice.ListTranscriptionJobsInput{
		JobNameContains: aws.String(jobName),
	}
	res, err := svc.ListTranscriptionJobs(input)
	if err != nil {
		return "", err
	}
	return aws.StringValue(res.TranscriptionJobSummaries[0].TranscriptionJobStatus), nil
}
