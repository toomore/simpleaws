package sqs

import (
	"fmt"
	"os"
	"testing"
)

var URL = "https://sqs.ap-northeast-1.amazonaws.com/271756324461/test_toomore"
var sqsqueue = New(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "ap-northeast-1", URL)

func TestSend(t *testing.T) {
	fmt.Println(sqsqueue.Send("Send msg to sqs from simplesqs."))
}

func TestReceive(t *testing.T) {
	if rece, err := sqsqueue.Receive(30); err == nil {
		for i, msg := range rece.Messages {
			t.Logf("[%d] %s [%s]", i, *msg.MessageID, *msg.Body)
		}
	}
}
