package s3

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var object = NewObject("toomore-aet", "simples3-test.txt", bytes.NewReader([]byte("This is test for simples3.")))
	var s3bucket = New(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "us-east-1")
	fmt.Println(s3bucket.PutObject(object))
}
