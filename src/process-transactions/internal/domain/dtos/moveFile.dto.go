package dtos

import "jucabet/stori-challenge/process-transactions/internal/domain/enums"

type MoveFileDto struct {
	Folder   enums.StorageFolders
	Filename string
}
