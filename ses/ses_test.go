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
	ses := New(os.Getenv("AWSID"), os.Getenv("AWSKEY"))
	msg := Message([]*mail.Address{user}, sender, "This is send from SimpleSES.", "<b>Hello Toomore~.</b>")
	result, err := ses.SendEmail(msg)
	t.Logf("%+v [Error: %s]\n", result, err)
}
