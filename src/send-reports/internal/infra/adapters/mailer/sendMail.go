package mailer

import (
	"gopkg.in/gomail.v2"

	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
)

func (adapter *Mailer) SendEmail(message *dtos.SendEmailDto) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", adapter.fromEmail)
	msg.SetHeader("To", message.Email)
	msg.SetHeader("Subject", message.Subject)
	msg.SetBody("text/html", message.Content)

	// Send the email to Bob
	dialer := gomail.NewDialer(adapter.smtpHost, adapter.smtpPort, adapter.user, adapter.password)
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
