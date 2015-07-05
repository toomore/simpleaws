package ses

import (
	"net/mail"
	"os"
	"testing"
)

var user = &mail.Address{
	Name:    "講太多",
	Address: "toomore0929@gmail.com",
}

var sender = &mail.Address{
	Name:    "講太多",
	Address: "me@toomore.net",
}

func TestNew(t *testing.T) {
	ses := New(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "us-east-1")
	result, err := ses.Send(sender, []*mail.Address{user}, "This is send from SimpleSES.", "<b>Hello Toomore~.</b>")
	t.Logf("%+v [Error: %s]\n", result, err)
}
