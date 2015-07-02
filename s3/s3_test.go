package s3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awsutil"
)

var s3bucket = New(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "us-east-1")

func TestNewPutObject(t *testing.T) {
	var object = NewPutObject("toomore-aet", "simples3-test.txt", bytes.NewReader([]byte("This is test for simples3.")))
	resp, _ := s3bucket.PutObject(object)
	fmt.Println(awsutil.StringValue(resp))
}

func TestNewGetObject(t *testing.T) {
	var object = NewGetObject("toomore-aet", "simples3-test.txt")
	resp, _ := s3bucket.GetObject(object)
	fmt.Println(awsutil.StringValue(resp))
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("[%s]\n", body)
	}
}
