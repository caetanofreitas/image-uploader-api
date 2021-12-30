package environment

import (
	"os"
	"strings"
)

var DATABASE_TYPE string
var DATABASE_CONNECTION string
var UPLOAD_URL string
var USE_AWS string
var AWS_REGION string
var AWS_ACCESS_KEY string
var AWS_SECRET string
var S3_NAME string
var S3_URL string

func failPanic(env string, name string) {
	if strings.Compare(env, "") == 0 {
		panic("Failed to load " + name + " env")
	}
}

func ValidateEnvironment() {
	DATABASE_TYPE = os.Getenv("DATABASE_TYPE")
	failPanic(DATABASE_TYPE, "DATABASE_TYPE")
	DATABASE_CONNECTION = os.Getenv("DATABASE_CONNECTION")
	failPanic(DATABASE_CONNECTION, "DATABASE_CONNECTION")
	UPLOAD_URL = os.Getenv("UPLOAD_URL")

	USE_AWS = os.Getenv("USE_AWS")
	if strings.Compare(USE_AWS, "") != 0 {
		AWS_REGION = os.Getenv("AWS_REGION")
		failPanic(AWS_REGION, "AWS_REGION")
		AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
		failPanic(AWS_ACCESS_KEY, "AWS_ACCESS_KEY")
		AWS_SECRET = os.Getenv("AWS_SECRET")
		failPanic(AWS_SECRET, "AWS_SECRET")
		S3_NAME = os.Getenv("S3_NAME")
		failPanic(S3_NAME, "S3_NAME")
		S3_URL = os.Getenv("S3_URL")
		failPanic(S3_URL, "S3_URL")
	}
}
