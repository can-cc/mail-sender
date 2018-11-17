package main

import (
	"github.com/fwchen/mail-sender/email"
	"github.com/fwchen/mail-sender/email/mailgun"
	"github.com/fwchen/mail-sender/web"
)

func main() {
	client := web.NewHTTPClient()
	sender := mailgun.NewSender(client, "xxx", "0cc9d726")
	to := email.Recipient{
		Name:    "Me",
		Address: "xxx",
	}
	sender.Send("xxx@xx", email.Params{
		"name": "Hello",
	}, "Test", to)
}
