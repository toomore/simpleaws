// ses - simple to send mails.
package ses

import (
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ses"
)

// New to new a ses.SES
func New(AWSID, AWSKEY string) *ses.SES {
	var config = aws.DefaultConfig
	config.Region = "us-east-1"
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return ses.New(config)
}

// Message to render a ses.SendEmailInput
func Message(ToUsers []*mail.Address, Sender *mail.Address, Subject,
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
