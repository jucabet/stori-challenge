package ports

import (
	"jucabet/stori-challenge/process-transactions/internal/domain/dtos"
	"jucabet/stori-challenge/process-transactions/internal/domain/enums"
)

type IStorageService interface {
	GetFilesByFolder(enums.StorageFolders) ([]string, error)
	GetFileContentByName(folder enums.StorageFolders, name string) ([]*dtos.TransactionsFileInfoDto, error)
}
