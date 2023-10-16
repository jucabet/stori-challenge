package factory

import (
	"jucabet/stori-challenge/send-reports/internal/infra/adapters/mailer"
	"os"

	"github.com/mailjet/mailjet-apiv3-go/v3"
)

func NewMailerAdapter() (*mailer.Mailer, error) {
	client := mailjet.NewMailjetClient(os.Getenv("MAILER_SERVICE_PUBLIC_KEY"), os.Getenv("MAILER_SERVICE_SECRET_KEY"))
	return mailer.NewMailer(
		client,
		os.Getenv("MAILER_FROM_EMAIL"),
		os.Getenv("MAILER_FROM_USER"),
	), nil
}
