package bucket

import "aws-sdk-go/pkg/store"

type API interface {
	CreateBucket(in store.BucketIn) (error, *store.CreateBucketOutput)
	DeleteBucket(in store.BucketIn) (error, *store.DeleteBucketOutput)
	ListObjects(in store.BucketIn) (error, *store.ListObjectsOutput)
}
