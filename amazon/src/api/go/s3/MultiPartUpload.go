package bucket

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func Upload(s3Client *s3.Client, resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNum int) (completedPart types.CompletedPart, err error) {
	var try int
	RETRIES := 2
	for try <= RETRIES {
		uploadResp, err := s3Client.UploadPart(context.TODO(), &s3.UploadPartInput{
			Body:          bytes.NewReader(fileBytes),
			Bucket:        resp.Bucket,
			Key:           resp.Key,
			PartNumber:    int32(partNum),
			UploadId:      resp.UploadId,
			ContentLength: int64(len(fileBytes)),
		})

		// Upload failed
		if err != nil {
			fmt.Println(err)
			// Max retries reached! Quitting
			if try == RETRIES {
				return types.CompletedPart{}, err
			} else {
				// Retrying
				try++
			}
		} else {
			// Upload is done!
			return types.CompletedPart{
				ETag:       uploadResp.ETag,
				PartNumber: int32(partNum),
			}, nil
		}
	}
	return types.CompletedPart{}, nil
}
