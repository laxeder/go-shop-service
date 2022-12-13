package user

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type UserDocument struct {
	Document    string `json:"document,omitempty" redis:"document,omitempty"`
	OldDocument string `json:"old_document,omitempty" redis:"old_document,omitempty"`
}

func NewUserDocument(userDocumentByte ...[]byte) (userDocument *UserDocument, err error) {
	var log = logger.New()

	userDocument = &UserDocument{}
	err = nil

	if len(userDocumentByte) == 0 {
		return userDocument, err
	}

	err = json.Unmarshal(userDocumentByte[0], userDocument)
	if err != nil {
		log.Error().Err(err).Msgf("O json do usuário está incorreto. %v", userDocumentByte[0])
		return userDocument, err
	}

	return userDocument, err
}
