package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQS struct {
	url *string
	sqs *sqs.SQS
}

func (s SQS) Send(Body string) (*sqs.SendMessageOutput, error) {
	return s.sqs.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(Body),
		QueueURL:    s.url,
	})
}

func (s SQS) Receive(Visibility int64) (*sqs.ReceiveMessageOutput, error) {
	return s.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Long(10),
		QueueURL:            s.url,
		VisibilityTimeout:   aws.Long(Visibility),
	})
}

// New to new a ses.S3
func New(AWSID, AWSKEY, Region, URL string) *SQS {
	var config = aws.DefaultConfig
	config.Region = Region
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return &SQS{
		url: aws.String(URL),
		sqs: sqs.New(config),
	}
}
