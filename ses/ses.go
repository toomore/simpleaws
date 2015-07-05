// Package ses - simple for ses.
package ses

import (
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SES struct
type SES struct {
	ses    *ses.SES
	Sender *mail.Address
}

// New to new a ses
func New(AWSID, AWSKEY, Region string) *SES {
	var config = aws.DefaultConfig
	config.Region = Region
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return &SES{
		ses: ses.New(config),
	}
}

// Send to send mail.
func (s SES) Send(Sender *mail.Address, ToUsers []*mail.Address, Subject,
	Content string) (*ses.SendEmailOutput, error) {
	return s.ses.SendEmail(Message(Sender, ToUsers, Subject, Content))
}

// Message to render a ses.SendEmailInput
func Message(Sender *mail.Address, ToUsers []*mail.Address, Subject,
	Content string) *ses.SendEmailInput {

	var mailCharset = aws.String("UTF-8")
	var toUsers []*string

	toUsers = make([]*string, len(ToUsers))
	for i, v := range ToUsers {
		toUsers[i] = aws.String(v.String())
	}

	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toUsers,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				HTML: &ses.Content{
					Charset: mailCharset,
					Data:    aws.String(Content),
				},
			},
			Subject: &ses.Content{
				Charset: mailCharset,
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender.String()),
	}
}
