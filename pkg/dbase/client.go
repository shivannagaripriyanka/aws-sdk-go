package dbase

import "aws-sdk-go/pkg/store"

type API interface {
	BatchGetItem(in *store.BatchGetItemInput) (error, *store.BatchGetItemOutput)
	BatchWriteItem(in *store.BatchWriteItemInput) (error, *store.BatchWriteItemOutput)
	//Close() error
}
