package application

import (
	"jucabet/stori-challenge/process-transactions/mocks"
	"testing"
)

func TestNewProcessTransactions(t *testing.T) {
	storageService := new(mocks.IStorageService)
	storageRepo := new(mocks.IStorageRepository)
	databaseRepo := new(mocks.IDatabaseRepository)
	topicRepo := new(mocks.ITopicRepository)

	usecase := NewProcessTransactions(
		storageService,
		storageRepo,
		databaseRepo,
		topicRepo,
	)

	if usecase == nil {
		t.Fatal("Error in NewProcessTransactions")
	}
}
