package s3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

var object = &s3.PutObjectInput{
	Bucket: aws.String("toomore-aet"),
	Key:    aws.String("simples3-test.txt"),
	Body:   bytes.NewReader([]byte("This is test for simples3.")),
}

// New to new a ses.SES
func New(AWSID, AWSKEY, Region string) *s3.S3 {
	var config = aws.DefaultConfig
	config.Region = Region
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return s3.New(config)
}
