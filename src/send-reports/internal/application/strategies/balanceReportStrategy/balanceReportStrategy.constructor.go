package balancereportstrategy

import (
	"jucabet/stori-challenge/send-reports/internal/application/interfaces"
	"jucabet/stori-challenge/send-reports/internal/domain/ports"
)

type BalanceReportStrategy struct {
	databaseService ports.IDatabaseService
	mailerService   ports.IMailerService
}

func NewBalanceReportStrategy(
	databaseService ports.IDatabaseService,
	mailerService ports.IMailerService,
) interfaces.IStategiesInterface {
	return &BalanceReportStrategy{
		databaseService: databaseService,
		mailerService:   mailerService,
	}
}
