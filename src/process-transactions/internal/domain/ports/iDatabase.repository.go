package ports

import "jucabet/stori-challenge/process-transactions/internal/domain/entities"

type IDatabaseRepository interface {
	SaveFileInfo(*entities.FileCharge) error
	SaveTransaction(*entities.Transaction) error
}
