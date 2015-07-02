package s3

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var s3bucket = New(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "us-east-1")
	fmt.Println(s3bucket.PutObject(object))
}
