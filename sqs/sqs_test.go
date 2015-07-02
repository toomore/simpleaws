package sqs

import (
	"fmt"
	"os"
	"testing"
)

var sqsqueue = New(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "ap-northeast-1", "https://sqs.ap-northeast-1.amazonaws.com/271756324461/test_toomore")

func TestNew(t *testing.T) {
	fmt.Println(sqsqueue.Send("Send msg to sqs from simplesqs."))
}
