package Utilities

import (
	"aws-sdk-go/pkg/store"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

const InvalidBucketName = "InvalidBucketName or EmptyBucketName"
const InvalidTableName = "InvalidTableName or EmptyTableName"

var Region = os.Getenv("AWS_REGION")

var envErr = godotenv.Load(".env")

func BadRequest(errorf string, reqData store.BucketIn) {
	//c.AbortWithError(http.StatusBadRequest, fmt.Errorf(utilities.BadRequest400))
	LogInfoWithErrorf("END", reqData.BucketName, "CreateBucket", errorf)
}
func BadReq(errorf string, reqData store.TableIn) {
	//c.AbortWithError(http.StatusBadRequest, fmt.Errorf(utilities.BadRequest400))
	LogInfoWithErrorf("END", reqData.TableName, "CreateBucket", errorf)
}
func LogInfoWithErrorf(msg string, buckename string, methodName string, err_msg string) {

	logrus.WithFields(logrus.Fields{
		"[BucketName]": buckename,
		"MethodName":   methodName,
		"Error":        err_msg,
	}).Info(msg)
}
