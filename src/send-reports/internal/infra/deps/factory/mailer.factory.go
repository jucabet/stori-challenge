package factory

import (
	"jucabet/stori-challenge/send-reports/internal/infra/adapters/mailer"
	"os"
	"strconv"
)

func NewMailerAdapter() (*mailer.Mailer, error) {
	smtpPort, _ := strconv.Atoi(os.Getenv("MAILER_FROM_PORT"))
	return mailer.NewMailer(
		os.Getenv("MAILER_FROM_HOST"),
		smtpPort,
		os.Getenv("MAILER_FROM_EMAIL"),
		os.Getenv("MAILER_FROM_USER"),
		os.Getenv("MAILER_FROM_PASSWORD"),
	), nil
}
