package product

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type ProductUid struct {
	Uid    string `json:"document,omitempty" redis:"document,omitempty"`
	OldUid string `json:"old_document,omitempty" redis:"old_document,omitempty"`
}

func NewProductUid(productUidByte ...[]byte) (productUid *ProductUid, err error) {
	var log = logger.New()

	productUid = &ProductUid{}
	err = nil

	if len(productUidByte) == 0 {
		return productUid, err
	}

	err = json.Unmarshal(productUidByte[0], productUid)
	if err != nil {
		log.Error().Err(err).Msgf("O json do produto está incorreto. %v", productUidByte[0])
		return productUid, err
	}

	return productUid, err
}
