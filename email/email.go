package email

import (
	// "bytes"
	// "html/template"
	"net/mail"
	// "regexp"
	// "strings"
	// "github.com/fwchen/mail-sender/models"
	// "github.com/getfider/fider/app/pkg/env"
)

// var cache = make(map[string]*template.Template, 0)

// Params used to replace variables on emails
type Params map[string]interface{}

// Merge given params into current params
func (p Params) Merge(p2 Params) Params {
	for k, v := range p2 {
		p[k] = v
	}
	return p
}

// Message represents what is sent by email
type Message struct {
	Subject string
	Body    string
}

// var baseTpl, _ = template.ParseFiles(env.Path("/views/templates/base_email.tpl"))

// RenderMessage returns the HTML of an email based on template and params
func RenderMessage(subject string, body string) *Message {

	return &Message{
		Subject: subject,
		Body:    body,
	}
}

// Context holds everything emailers need to know about execution context
type Context interface {
	BaseURL() string
	LogoURL() string
}

// NoReply is the default 'from' address
// var NoReply = env.MustGet("EMAIL_NOREPLY")

// Recipient contains details of who is receiving the email
type Recipient struct {
	Name    string
	Address string
	Params  Params
}

// NewRecipient creates a new Recipient
func NewRecipient(name, address string, params Params) Recipient {
	return Recipient{
		Name:    name,
		Address: address,
		Params:  params,
	}
}

// Strings returns the RFC format to send emails via SMTP
func (r Recipient) String() string {
	if r.Address == "" {
		return ""
	}

	address := mail.Address{
		Name:    r.Name,
		Address: r.Address,
	}

	return address.String()
}

// var whitelist = env.GetEnvOrDefault("EMAIL_WHITELIST", "")
// var whitelistRegex = regexp.MustCompile(whitelist)
// var blacklist = env.GetEnvOrDefault("EMAIL_BLACKLIST", "")
// var blacklistRegex = regexp.MustCompile(blacklist)

// SetWhitelist can be used to change email whitelist during runtime
// func SetWhitelist(s string) {
// 	whitelist = s
// 	whitelistRegex = regexp.MustCompile(whitelist)
// }

// // SetBlacklist can be used to change email blacklist during runtime
// func SetBlacklist(s string) {
// 	blacklist = s
// 	blacklistRegex = regexp.MustCompile(blacklist)
// }

// // CanSendTo returns true if Fider is allowed to send email to given address
// func CanSendTo(address string) bool {
// 	if strings.TrimSpace(address) == "" {
// 		return false
// 	}

// 	if whitelist != "" {
// 		return whitelistRegex.MatchString(address)
// 	}

// 	if blacklist != "" {
// 		return !blacklistRegex.MatchString(address)
// 	}

// 	return true
// }

// Sender is used to send emails
type Sender interface {
	Send(ctx Context, templateName string, params Params, from string, to Recipient) error
	BatchSend(ctx Context, templateName string, params Params, from string, to []Recipient) error
}
