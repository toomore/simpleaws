// Package sqs - simple for sqs.
package sqs

import (
	"runtime"
	"sync"

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

func (s SQS) SendBatch(Bodies []string) (*sqs.SendMessageBatchOutput, error) {
	var entries []*sqs.SendMessageBatchRequestEntry
	entries = make([]*sqs.SendMessageBatchRequestEntry, len(Bodies))

	for i, body := range Bodies {
		entries[i] = &sqs.SendMessageBatchRequestEntry{
			ID:          aws.String(string(97 + i)),
			MessageBody: aws.String(body),
		}
	}
	return s.sqs.SendMessageBatch(&sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueURL: s.url,
	})
}

type Batch struct {
	Output *sqs.SendMessageBatchOutput
	Error  error
}

func (s SQS) SendBatchList(Bodies []string) []*Batch {
	var (
		BodiesLen   = len(Bodies)
		maxlen      = 10
		times       = BodiesLen / maxlen
		more        = BodiesLen % maxlen
		wg          sync.WaitGroup
		result      chan *Batch
		do          func([]string)
		BatchOutput []*Batch
	)

	result = make(chan *Batch)
	do = func(Bodies []string) {
		defer wg.Done()
		runtime.Gosched()
		var b = &Batch{}
		b.Output, b.Error = s.SendBatch(Bodies)
		result <- b
	}

	wg.Add(times)

	if more > 0 {
		wg.Add(1)
		go do(Bodies[maxlen*times : maxlen*times+more])
	}
	for i := 0; i < times; i++ {
		go do(Bodies[maxlen*i : maxlen*(i+1)])
	}

	BatchOutput = make([]*Batch, 0)
	go func() {
		for {
			select {
			case v, ok := <-result:
				if ok {
					BatchOutput = append(BatchOutput, v)
				}
			}
		}
	}()
	wg.Wait()
	return BatchOutput
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
