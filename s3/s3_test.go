package s3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awsutil"
)

var s3bucket = New(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "us-east-1", "toomore-aet")

func TestNewPutObject(t *testing.T) {
	resp, _ := s3bucket.Put("simples3-test.txt", bytes.NewReader([]byte("This is test for simples3.")))
	fmt.Println(awsutil.StringValue(resp))
}

func TestNewGetObject(t *testing.T) {
	resp, _ := s3bucket.Get("simples3-test.txt")
	fmt.Println(awsutil.StringValue(resp))
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Printf("[%s]\n", body)
	}
}
