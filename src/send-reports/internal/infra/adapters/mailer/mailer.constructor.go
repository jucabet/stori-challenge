package mailer

import "github.com/mailjet/mailjet-apiv3-go/v3"

type Mailer struct {
	client    *mailjet.Client
	fromEmail string
	user      string
}

func NewMailer(
	client *mailjet.Client,
	fromEmail string,
	user string,
) *Mailer {
	return &Mailer{
		client:    client,
		fromEmail: fromEmail,
		user:      user,
	}
}
