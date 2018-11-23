package main

import (
	"flag"
	"github.com/fwchen/mail-sender/email"
	"github.com/fwchen/mail-sender/email/mailgun"
	"github.com/fwchen/mail-sender/web"
)

func main() {
	mailgunApi := flag.String("mailgun-api", "", "mailgun api")
	domain := flag.String("domain", "", "mailgun domain")
	noReply := flag.String("noreply", "", "mailgun noreply")
	recipient := flag.String("recipient", "", "recipient address")
	recipientName := flag.String("target-name", "", "recipient name")
	flag.Parse()

	if len(*mailgunApi) <= 0 {
		panic("mailgun-api is required.")
	}
	if len(*domain) <= 0 {
		panic("domain is required.")
	}
	if len(*noReply) <= 0 {
		panic("noreply is required.")
	}
	if len(*recipient) <= 0 {
		panic("target is required.")
	}

	client := web.NewHTTPClient()
	sender := mailgun.NewSender(client, *domain, *mailgunApi)
	to := email.Recipient{
		Name:    *recipientName,
		Address: *recipient,
	}
	sender.Send(*noReply, email.Params{}, "Test", to)
}
