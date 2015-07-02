// Package s3 - simple for s3.
package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 struct.
type S3 struct {
	bucket   *string
	s3bucket *s3.S3
}

// Get to get a object from s3.
func (s S3) Get(Key string) (*s3.GetObjectOutput, error) {
	return s.s3bucket.GetObject(&s3.GetObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(Key),
	})
}

// Put to put a object into s3.
func (s S3) Put(Key string, Body io.ReadSeeker) (*s3.PutObjectOutput, error) {
	return s.s3bucket.PutObject(&s3.PutObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(Key),
		Body:   Body,
	})
}

// New to new a ses.S3
func New(AWSID, AWSKEY, Region, Bucket string) *S3 {
	var config = aws.DefaultConfig
	config.Region = Region
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return &S3{
		bucket:   aws.String(Bucket),
		s3bucket: s3.New(config),
	}
}
