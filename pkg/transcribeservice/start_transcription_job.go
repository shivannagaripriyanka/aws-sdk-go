package transcribeservice

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"strings"
	"time"

	//"github.com/aws/aws-sdk-go-v2/aws/client"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	//"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/smithy-go/middleware"
)

type TClient struct {
	*Client
}

type Client struct {
	request.Retryer
	metadata.ClientInfo

	Config   aws.Config
	Handlers request.Handlers
}

type StartTranscriptionJobInput struct {
	Media                     *types.Media
	TranscriptionJobName      *string
	ContentRedaction          *types.ContentRedaction
	IdentifyLanguage          *bool
	IdentifyMultipleLanguages *bool
	JobExecutionSettings      *types.JobExecutionSettings
	KMSEncryptionContext      map[string]string
	LanguageCode              types.LanguageCode
	LanguageIdSettings        map[string]types.LanguageIdSettings
	LanguageOptions           []types.LanguageCode

	// Specify the format of your input media file.
	MediaFormat              types.MediaFormat
	MediaSampleRateHertz     *int32
	ModelSettings            *types.ModelSettings
	OutputBucketName         *string
	OutputEncryptionKMSKeyId *string
	OutputKey                *string
	Settings                 *types.Settings
	Subtitles                *types.Subtitles
	Tags                     []types.Tag
}

type StartTranscriptionJobOutput struct {

	// Provides detailed information about the current transcription job, including job
	// status and, if applicable, failure reason.
	TranscriptionJob *types.TranscriptionJob
	ResultMetadata   middleware.Metadata
}

func (c *Client) StartTranscriptionJob(ctx context.Context, params *StartTranscriptionJobInput, optFns ...func(*Options)) (*StartTranscriptionJobOutput, error) {
	if params == nil {
		params = &StartTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTranscriptionJob", params, optFns, c.addOperationStartTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//func (c *Client) StartTranscriptionJob(input *StartTranscriptionJobInput) (*StartTranscriptionJobOutput, error) {
//	req, out := c.StartTranscriptionJobRequest(input)
//	return out, req.Send()
//}

func (c *TClient) StartTranscriptionJobRequest(input *StartTranscriptionJobInput) (req *request.Request, output *StartTranscriptionJobOutput) {
	op := &request.Operation{
		Name:       "opStartTranscriptionJob",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTranscriptionJobInput{}
	}

	output = &StartTranscriptionJobOutput{}
	req = c.Client.NewRequest(op, input, output)

	inputparm := &transcribeservice.GetTranscriptionJobInput{
		TranscriptionJobName: aws.String(jobName),
	}

	res, err := svc.GetTranscriptionJob(input)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return
}

func New(p client.ConfigProvider, cfgs ...*aws.Config) *TClient {
	//	// Create a TranscribeService client from just a session.
	svc := transcribeservice.New(mySession)
	//
	//	// Create a TranscribeService client with additional configuration
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "transcribe"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName, c.ResolvedRegion)
}

func startTranscription(filedata string) (string, error) {
	t := time.Now()
	b64data := filedata[strings.IndexByte(filedata, ',')+1:]
	data, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return "", err
	}
	sess, err := session.NewSession()
	if err != nil {
		return "", err
	}
	contentType := "audio/mp3"
	filename := t.Format(layout2) + ".mp3"
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(bucketName),
		Key:         aws.String(S3MediaPath + "/" + filename),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}
	url := "s3://" + bucketName + "/" + S3MediaPath + "/" + filename

	svc := transcribeservice.New(session.New())

	input := &transcribeservice.StartTranscriptionJobInput{
		TranscriptionJobName: aws.String(filename),
		LanguageCode:         aws.String(languageCode),
		OutputBucketName:     aws.String(bucketName),
		MediaFormat:          aws.String(mediaFormat),
		Media: &transcribeservice.Media{
			MediaFileUri: aws.String(url),
		},
	}
	_, err = svc.StartTranscriptionJob(input)
	if err != nil {
		return "", err
	}

	return filename, nil
}
