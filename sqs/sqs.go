// Package sqs - simple for sqs.
package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQS struct
type SQS struct {
	url *string
	sqs *sqs.SQS
}

// Send to send a queue message.
func (s SQS) Send(Body string) (*sqs.SendMessageOutput, error) {
	return s.sqs.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(Body),
		QueueURL:    s.url,
	})
}

// Receive to receive a queue message.
func (s SQS) Receive(Visibility int64) (*sqs.ReceiveMessageOutput, error) {
	return s.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Long(10),
		QueueURL:            s.url,
		VisibilityTimeout:   aws.Long(Visibility),
	})
}

// Delete to delete a queue message.
func (s SQS) Delete(ReceiptHandle *string) (*sqs.DeleteMessageOutput, error) {
	return s.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueURL:      s.url,
		ReceiptHandle: ReceiptHandle,
	})
}

// New to new a sqs.
func New(AWSID, AWSKEY, Region, URL string) *SQS {
	var config = aws.DefaultConfig
	config.Region = Region
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return &SQS{
		url: aws.String(URL),
		sqs: sqs.New(config),
	}
}
