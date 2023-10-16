package application

import "jucabet/stori-challenge/send-reports/internal/domain/ports"

type SendReport struct {
	databaseService ports.IDatabaseService
	mailerService   ports.IMailerService
}

func NewSendReport(
	databaseService ports.IDatabaseService,
	mailerService ports.IMailerService,
) *SendReport {
	return &SendReport{
		databaseService: databaseService,
		mailerService:   mailerService,
	}
}
