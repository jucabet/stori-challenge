package ports

import "jucabet/stori-challenge/send-reports/internal/domain/dtos"

type IMailerService interface {
	SendEmail(*dtos.SendEmailDto) error
}
