package mailgun

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/fwchen/mail-sender/email"
	// "github.com/getfider/fider/app/pkg/env"
	// "github.com/getfider/fider/app/pkg/errors"
	// "github.com/getfider/fider/app/pkg/log"
	"github.com/fwchen/mail-sender/errors"
	"github.com/fwchen/mail-sender/web"
)

var baseURL = "https://api.mailgun.net/v3/%s/messages"

type Props map[string]interface{}

//Sender is used to send emails
type Sender struct {
	// logger log.Logger
	client web.Client
	domain string
	apiKey string
}

//NewSender creates a new mailgun email sender
func NewSender(client web.Client, domain, apiKey string) *Sender {
	return &Sender{client, domain, apiKey}
}

//Send an email
func (s *Sender) Send(noReply string, params email.Params, from string, to email.Recipient) error {
	return s.BatchSend(noReply, params, from, []email.Recipient{to})
}

// BatchSend an email to multiple recipients
func (s *Sender) BatchSend(noReply string, params email.Params, from string, to []email.Recipient) error {
	if len(to) == 0 {
		return nil
	}

	isBatch := len(to) > 1

	var message *email.Message
	if isBatch {
		// Replace recipient specific Go templates variables with Mailgun template variables
		for k := range to[0].Params {
			params[k] = fmt.Sprintf("%%recipient.%s%%", k)
		}
		message = email.RenderMessage("subject", "body")
	} else {
		message = email.RenderMessage("subject", "body")
	}

	form := url.Values{}
	form.Add("from", email.NewRecipient(from, noReply, email.Params{}).String())
	form.Add("h:Reply-To", noReply)
	form.Add("subject", message.Subject)
	form.Add("html", message.Body)

	// Set Mailgun's var based on each recipient's variables
	recipientVariables := make(map[string]email.Params, 0)
	for _, r := range to {
		if r.Address != "" {
			form.Add("to", r.String())
			recipientVariables[r.Address] = r.Params
		}
	}

	// If we skipped all recipients, just return
	if len(recipientVariables) == 0 {
		return nil
	}

	if isBatch {
		json, err := json.Marshal(recipientVariables)
		if err != nil {
			return errors.Wrap(err, "failed to marshal recipient variables")
		}

		form.Add("recipient-variables", string(json))
	}

	if isBatch {
		fmt.Println("Sending email to @{CountRecipients} recipients.", Props{
			"CountRecipients": len(recipientVariables),
		})
	} else {
		fmt.Println("Sending email to @{Address}.", map[string]string{
			"Address": to[0].Address,
		})
	}

	url := fmt.Sprintf(baseURL, s.domain)
	request, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		return errors.Wrap(err, "failed to create POST request")
	}

	request.SetBasicAuth("api", s.apiKey)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}

	defer resp.Body.Close()
	fmt.Println("Email sent with response code @{StatusCode}.", Props{
		"StatusCode": resp.StatusCode,
	})
	return nil
}
