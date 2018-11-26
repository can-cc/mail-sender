package main

import (
	"flag"
	"fmt"
	"github.com/fwchen/mail-sender/email"
	"github.com/fwchen/mail-sender/email/mailgun"
	"github.com/fwchen/mail-sender/web"
	"io/ioutil"
)

func main() {
	flag.Usage = func() {
		b, err := ioutil.ReadFile("help.txt") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		helpStr := string(b)
		fmt.Println(helpStr)
	}

	mailgunApi := flag.String("mailgun-api", "", "mailgun api")
	domain := flag.String("domain", "", "mailgun domain")
	noReply := flag.String("noreply", "", "mailgun noreply")
	recipient := flag.String("recipient", "", "recipient address")
	recipientName := flag.String("recipient-name", "", "recipient name")
	flag.Parse()

	if len(*mailgunApi) <= 0 {
		flag.Usage()
		return
	}
	if len(*domain) <= 0 {
		flag.Usage()
		return
	}
	if len(*noReply) <= 0 {
		flag.Usage()
		return
	}
	if len(*recipient) <= 0 {
		flag.Usage()
		return
	}
	if len(*recipientName) <= 0 {
		flag.Usage()
		return
	}

	fmt.Println()

	client := web.NewHTTPClient()
	sender := mailgun.NewSender(client, *domain, *mailgunApi)
	to := email.Recipient{
		Name:    *recipientName,
		Address: *recipient,
	}
	sender.Send(*noReply, email.Params{}, "Test", to)

	fmt.Println()
}
