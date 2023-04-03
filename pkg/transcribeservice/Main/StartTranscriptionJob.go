package transcribeservice

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"strings"
	"time"
)

func StartTranscription(filedata string) (string, error) {
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
	filename := t.Format("layout2") + ".mp3"
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String("aws-sdk-go"),
		Key:         aws.String(S3MediaPath + "/" + filename),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}
	url := "s3://" + "aws-sdk-go" + "/" + S3MediaPath + "/" + filename

	svc := transcribeservice.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

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
