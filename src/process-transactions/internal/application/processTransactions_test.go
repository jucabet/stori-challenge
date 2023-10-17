package application

import (
	"errors"
	"jucabet/stori-challenge/process-transactions/internal/domain/dtos"
	"jucabet/stori-challenge/process-transactions/internal/domain/entities"
	"jucabet/stori-challenge/process-transactions/internal/domain/enums"
	"jucabet/stori-challenge/process-transactions/internal/domain/utils"
	"jucabet/stori-challenge/process-transactions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcessTransactions(t *testing.T) {
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

	t.Run("error getting files by folder", func(t *testing.T) {
		storageService.On("GetFilesByFolder", enums.IncomingFiles).
			Return([]string{}, errors.New("MockGetFilesByFolder")).
			Once()

		err := usecase.ProcessTransactions()
		assert.Error(t, err, "MockGetFilesByFolder")
	})

	t.Run("files not found", func(t *testing.T) {
		storageService.On("GetFilesByFolder", enums.IncomingFiles).
			Return([]string{}, nil).
			Once()

		err := usecase.ProcessTransactions()
		assert.Error(t, err, "FilesNotFound")
	})

	t.Run("error processing files", func(t *testing.T) {
		storageService.On("GetFilesByFolder", enums.IncomingFiles).
			Return([]string{"transactions.csv"}, nil).
			Once()

		mockProcessFiles := processFiles
		processFiles = func(pt *ProcessTransaction, files []string) error {
			return errors.New("mockProcessFiles")
		}

		err := usecase.ProcessTransactions()
		assert.Error(t, err, "mockProcessFiles")

		processFiles = mockProcessFiles
	})

	t.Run("sucess", func(t *testing.T) {
		storageService.On("GetFilesByFolder", enums.IncomingFiles).
			Return([]string{"transactions.csv"}, nil).
			Once()

		mockProcessFiles := processFiles
		processFiles = func(pt *ProcessTransaction, files []string) error {
			return nil
		}

		err := usecase.ProcessTransactions()
		assert.NoError(t, err)

		processFiles = mockProcessFiles
	})
}

func TestProcessFiles(t *testing.T) {
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

	files := []string{"transactions.csv"}
	txFileInfo := []*dtos.TransactionsFileInfoDto{
		{
			ID:          1,
			Date:        time.Now(),
			Transaction: 100000.10,
		},
	}

	mockProcessFileTransactions := processFileTransactions
	mockMoveFile := moveFile

	t.Run("error getting file content", func(t *testing.T) {
		storageService.On("GetFileContentByName", enums.IncomingFiles, files[0]).
			Return(nil, errors.New("mockGetFileContentByName")).
			Once()

		err := processFiles(usecase, files)
		assert.Error(t, err, "mockGetFileContentByName")
	})

	t.Run("error processing file transaction", func(t *testing.T) {
		storageService.On("GetFileContentByName", enums.IncomingFiles, files[0]).
			Return(txFileInfo, nil).
			Once()

		processFileTransactions = func(
			pt *ProcessTransaction,
			filename string,
			txFileInfo []*dtos.TransactionsFileInfoDto,
		) error {
			return errors.New("mockErrorProcessFileTransactions")
		}

		err := processFiles(usecase, files)
		assert.NoError(t, err)
	})

	t.Run("error move file", func(t *testing.T) {
		storageService.On("GetFileContentByName", enums.IncomingFiles, files[0]).
			Return(txFileInfo, nil).
			Once()

		processFileTransactions = func(
			pt *ProcessTransaction,
			filename string,
			txFileInfo []*dtos.TransactionsFileInfoDto,
		) error {
			return nil
		}

		moveFile = func(pt *ProcessTransaction, filename string) error {
			return errors.New("mockErrorMoveFile")
		}

		err := processFiles(usecase, files)
		assert.NoError(t, err)
	})

	t.Run("sucess", func(t *testing.T) {
		storageService.On("GetFileContentByName", enums.IncomingFiles, files[0]).
			Return(txFileInfo, nil).
			Once()

		processFileTransactions = func(
			pt *ProcessTransaction,
			filename string,
			txFileInfo []*dtos.TransactionsFileInfoDto,
		) error {
			return nil
		}

		moveFile = func(pt *ProcessTransaction, filename string) error {
			return nil
		}

		err := processFiles(usecase, files)
		assert.NoError(t, err)
	})

	processFileTransactions = mockProcessFileTransactions
	moveFile = mockMoveFile
}

func TestProcessFileTransactions(t *testing.T) {
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

	filename := "transaction.csv"
	fileChargeID := "abc123"
	txFileInfo := []*dtos.TransactionsFileInfoDto{
		{
			ID:          1,
			Date:        time.Now(),
			Transaction: 100000.10,
		},
	}

	mockRegisterFileChargeInfo := registerFileChargeInfo
	mockRegisterTransaction := registerTransaction

	t.Run("error register file charge info", func(t *testing.T) {
		registerFileChargeInfo = func(pt *ProcessTransaction, filename string) (string, error) {
			return "", errors.New("mockErrorRegisterFileChargeInfo")
		}

		err := processFileTransactions(usecase, filename, txFileInfo)
		assert.Error(t, err, "mockErrorRegisterFileChargeInfo")
	})

	t.Run("error register transaction and send message to queue", func(t *testing.T) {
		registerFileChargeInfo = func(pt *ProcessTransaction, filename string) (string, error) {
			return fileChargeID, nil
		}

		registerTransaction = func(
			pt *ProcessTransaction,
			fileChargeID string,
			transaction *dtos.TransactionsFileInfoDto,
		) error {
			return errors.New("mockErrorRegisterTransaction")
		}

		topicRepo.On("SendMessageToReport", fileChargeID).
			Return(errors.New("mockErrorSendMessageToReport")).
			Once()

		err := processFileTransactions(usecase, filename, txFileInfo)
		assert.Error(t, err, "mockErrorSendMessageToReport")
	})

	t.Run("sucess", func(t *testing.T) {
		registerFileChargeInfo = func(pt *ProcessTransaction, filename string) (string, error) {
			return fileChargeID, nil
		}

		registerTransaction = func(
			pt *ProcessTransaction,
			fileChargeID string,
			transaction *dtos.TransactionsFileInfoDto,
		) error {
			return nil
		}

		topicRepo.On("SendMessageToReport", fileChargeID).
			Return(nil).
			Once()

		err := processFileTransactions(usecase, filename, txFileInfo)
		assert.NoError(t, err)
	})

	registerFileChargeInfo = mockRegisterFileChargeInfo
	registerTransaction = mockRegisterTransaction
}

func TestRegisterFileChargeInfo(t *testing.T) {
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

	filename := "transaction.csv"
	fileChargeID := "123abc"
	now := time.Now()

	chargeInfo := &entities.FileCharge{
		Type:     enums.FileCharge,
		ID:       fileChargeID,
		Date:     now,
		FileName: filename,
	}

	mockGenerateRandomCode := utils.GenerateRandomCode
	mockGetCurrentTime := utils.GetCurrentTime
	utils.GenerateRandomCode = func(n int) string {
		return fileChargeID
	}

	utils.GetCurrentTime = func() time.Time {
		return now
	}

	t.Run("error save file info", func(t *testing.T) {
		databaseRepo.On("SaveFileInfo", chargeInfo).
			Return(errors.New("mockErrorSaveFileInfo")).
			Once()

		_, err := registerFileChargeInfo(usecase, filename)
		assert.Error(t, err, "mockErrorSaveFileInfo")
	})

	t.Run("sucess", func(t *testing.T) {
		databaseRepo.On("SaveFileInfo", chargeInfo).
			Return(nil).
			Once()

		fileChargeIDReceived, err := registerFileChargeInfo(usecase, filename)
		assert.NoError(t, err)
		assert.Equal(t, fileChargeIDReceived, fileChargeID)
	})

	utils.GenerateRandomCode = mockGenerateRandomCode
	utils.GetCurrentTime = mockGetCurrentTime
}

func TestRegisterTransaction(t *testing.T) {
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

	fileChargeID := "123abc"
	transaction := &dtos.TransactionsFileInfoDto{
		ID:          1,
		Date:        time.Now(),
		Transaction: 100000.10,
	}

	tx := &entities.Transaction{
		Type:             enums.Transaction,
		ID:               "1",
		Date:             transaction.Date,
		TransactionValue: transaction.Transaction,
		FileChargeID:     fileChargeID,
	}

	t.Run("error register database transaction", func(t *testing.T) {
		databaseRepo.On("SaveTransaction", tx).
			Return(errors.New("mockErrorSaveTransaction")).
			Once()

		err := registerTransaction(usecase, fileChargeID, transaction)
		assert.Error(t, err, "mockErrorSaveTransaction")
	})

	t.Run("sucess", func(t *testing.T) {
		databaseRepo.On("SaveTransaction", tx).
			Return(nil).
			Once()

		err := registerTransaction(usecase, fileChargeID, transaction)
		assert.NoError(t, err)
	})
}

func TestMoveFile(t *testing.T) {
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

	filename := "transactions.csv"

	src := &dtos.MoveFileDto{
		Folder:   enums.IncomingFiles,
		Filename: filename,
	}

	dest := &dtos.MoveFileDto{
		Folder:   enums.ProcessedFiles,
		Filename: filename,
	}

	t.Run("error run file", func(t *testing.T) {
		storageRepo.On("MoveFile", src, dest).
			Return(errors.New("mockErrorMoveFile")).
			Once()

		err := moveFile(usecase, filename)
		assert.Error(t, err, "mockErrorMoveFile")
	})

	t.Run("sucess", func(t *testing.T) {
		storageRepo.On("MoveFile", src, dest).
			Return(nil).
			Once()

		err := moveFile(usecase, filename)
		assert.NoError(t, err)
	})
}
