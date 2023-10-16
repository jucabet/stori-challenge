package ports

import "jucabet/stori-challenge/send-reports/internal/domain/entities"

type IDatabaseService interface {
	GetContacts() ([]*entities.Contact, error)
	GetTransactionsByFileChargeID(fileChargeID string) ([]*entities.Transaction, error)
}
