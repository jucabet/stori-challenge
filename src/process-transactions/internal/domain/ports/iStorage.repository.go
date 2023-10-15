package ports

import "jucabet/stori-challenge/process-transactions/internal/domain/dtos"

type IStorageRepository interface {
	MoveFile(src, dest *dtos.MoveFileDto) error
}
