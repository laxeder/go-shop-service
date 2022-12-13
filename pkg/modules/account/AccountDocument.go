package account

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type AccountDocument struct {
	Document    string `json:"document,omitempty" redis:"document,omitempty"`
	OldDocument string `json:"old_document,omitempty" redis:"old_document,omitempty"`
}

func NewAccountDocument(accountDocumentByte ...[]byte) (accountDocument *AccountDocument, err error) {
	var log = logger.New()

	accountDocument = &AccountDocument{}
	err = nil

	if len(accountDocumentByte) == 0 {
		return accountDocument, err
	}

	err = json.Unmarshal(accountDocumentByte[0], accountDocument)
	if err != nil {
		log.Error().Err(err).Msgf("O json do conta está incorreto. %v", accountDocumentByte[0])
		return accountDocument, err
	}

	return accountDocument, err
}
