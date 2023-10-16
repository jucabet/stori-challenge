package factory

import (
	localsecrets "jucabet/stori-challenge/send-reports/internal/infra/adapters/localSecrets"
	prodsecrets "jucabet/stori-challenge/send-reports/internal/infra/adapters/prodSecrets"
	"jucabet/stori-challenge/send-reports/internal/infra/enums"
	"jucabet/stori-challenge/send-reports/internal/infra/interfaces"
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
