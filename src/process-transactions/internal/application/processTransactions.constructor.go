package application

import "jucabet/stori-challenge/process-transactions/internal/domain/ports"

type ProcessTransaction struct {
	storageService ports.IStorageService
	storageRepo    ports.IStorageRepository
	databaseRepo   ports.IDatabaseRepository
	topicRepo      ports.ITopicRepository
}

func NewProcessTransactions(
	storageService ports.IStorageService,
	storageRepo ports.IStorageRepository,
	databaseRepo ports.IDatabaseRepository,
	topicRepo ports.ITopicRepository,
) *ProcessTransaction {
	return &ProcessTransaction{
		storageService: storageService,
		storageRepo:    storageRepo,
		databaseRepo:   databaseRepo,
		topicRepo:      topicRepo,
	}
}
