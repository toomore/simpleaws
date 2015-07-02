package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewObject(Bucket, Key string, Body io.ReadSeeker) *s3.PutObjectInput {
	return &s3.PutObjectInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(Key),
		Body:   Body,
	}
}

// New to new a ses.SES
func New(AWSID, AWSKEY, Region string) *s3.S3 {
	var config = aws.DefaultConfig
	config.Region = Region
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return s3.New(config)
}
