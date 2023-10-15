package factory

import (
	localsecrets "jucabet/stori-challenge/process-transactions/internal/infra/adapters/localSecrets"
	prodsecrets "jucabet/stori-challenge/process-transactions/internal/infra/adapters/prodSecrets"
	"jucabet/stori-challenge/process-transactions/internal/infra/enums"
	"jucabet/stori-challenge/process-transactions/internal/infra/interfaces"
	"os"
)

func NewSecretAdapter() interfaces.ISecretInterface {
	switch os.Getenv("ENV") {
	case string(enums.LOCAL):
		return localsecrets.NewLocalSecretAdapter()
	default:
		return prodsecrets.NewProdSecretAdapter()
	}
}
