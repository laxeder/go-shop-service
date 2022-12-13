package address

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type AddressDocument struct {
	Document    string `json:"document,omitempty" redis:"document,omitempty"`
	OldDocument string `json:"old_document,omitempty" redis:"old_document,omitempty"`
}

func NewAddressDocument(addressDocumentByte ...[]byte) (addressDocument *AddressDocument, err error) {
	var log = logger.New()

	addressDocument = &AddressDocument{}
	err = nil

	if len(addressDocumentByte) == 0 {
		return addressDocument, err
	}

	err = json.Unmarshal(addressDocumentByte[0], addressDocument)
	if err != nil {
		log.Error().Err(err).Msgf("O json do endereço está incorreto. %v", addressDocumentByte[0])
		return addressDocument, err
	}

	return addressDocument, err
}
