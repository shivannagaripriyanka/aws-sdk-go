package transcribe_tests

import (
	"aws-sdk-go/pkg/transcribeservice"
	"github.com/aws/aws-sdk-go/aws/client"
	"testing"
)

func TestStartTransciptionJobPass(t *testing.T) {
	out, err := transcribeservice.API.StartTranscriptionJob(client.Client{}, &transcribeservice.StartTranscriptionJobInput{})
}
