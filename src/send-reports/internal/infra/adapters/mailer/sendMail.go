package mailer

import (
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"

	"github.com/mailjet/mailjet-apiv3-go/v3"
)

func (adapter *Mailer) SendEmail(message *dtos.SendEmailDto) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: adapter.fromEmail,
				Name:  adapter.user,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: message.Email,
					Name:  message.Name,
				},
			},
			Subject:  message.Subject,
			HTMLPart: message.Content,
		},
	}

	mssgs := mailjet.MessagesV31{Info: messagesInfo}
	_, err := adapter.client.SendMailV31(&mssgs)
	if err != nil {
		return err
	}

	return nil
}
