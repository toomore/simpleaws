package sqs

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/toomore/simpleaws/utils"
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

func TestDelete(t *testing.T) {
	if rece, err := sqsqueue.Receive(30); err == nil {
		for i, msg := range rece.Messages {
			t.Logf("[%d] %s [%s]", i, *msg.MessageID, *msg.ReceiptHandle)
			t.Log(sqsqueue.Delete(msg.ReceiptHandle))
		}
	}
}

func TestSendBatch(t *testing.T) {
	sqsqueue.sendBatch([]string{"aa_1", "bb_2", "cc_3", "aa_4", "bb_5", "cc_6", "aa_7", "bb_8", "cc_9", "aa_10", "bb_11", "cc_12"})
}

func TestSendBatchList(t *testing.T) {
	sqsqueue.SendBatch([]string{"aa_1", "bb_2", "cc_3", "aa_4", "bb_5", "cc_6", "aa_7", "bb_8", "cc_9", "aa_10", "bb_11", "cc_12"})
}

func TestSendValuesData(t *testing.T) {
	m := map[string]string{"name": "Toomore", "age": "30"}
	vs := url.Values{}
	for k, v := range m {
		vs.Set(k, v)
	}
	t.Log(vs.Encode())
	sqsqueue.Send(vs.Encode())
	if resp, err := sqsqueue.Receive(10); err == nil {
		for _, msg := range resp.Messages {
			body, _ := utils.Base64Decode([]byte(*msg.Body))
			t.Log(url.ParseQuery(string(body)))
		}
	}
}
