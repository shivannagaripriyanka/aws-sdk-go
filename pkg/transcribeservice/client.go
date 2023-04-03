package transcribeservice

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
)

type Options struct {
}

type API interface {
	StartTranscriptionJobRequest(input *StartTranscriptionJobInput) (req *request.Request, output *StartTranscriptionJobOutput)
	StartTranscriptionJob(ctx context.Context, params *StartTranscriptionJobInput, optFns ...func(*Options)) (*StartTranscriptionJobOutput, error)
}