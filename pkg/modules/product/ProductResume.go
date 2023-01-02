package product

import (
	"encoding/json"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type ProductResume struct {
	ProductUid      string          `json:"product_uid,omitempty" redis:"product_uid,omitempty"`
	FreightUid      string          `json:"freight_uid,omitempty" redis:"freight_uid,omitempty"`
	Product         Product         `json:"product,omitempty" redis:"product,omitempty"`
	Freight         freight.Freight `json:"freight,omitempty" redis:"freight,omitempty"`
	ZipcodeReceiver string          `json:"zipcode_receiver,omitempty" redis:"zipcode_receiver,omitempty"`
}

func UnmarshalProductsResume(productResByte ...[]byte) (productResume *[]ProductResume, err error) {
	var log = logger.New()

	productResume = &[]ProductResume{}
	err = nil

	fmt.Print(len(productResByte))

	if len(productResByte) <= 0 {
		return
	}

	err = json.Unmarshal(productResByte[0], productResume)
	if err != nil {
		log.Error().Err(err).Msgf("O json do product resume está incorreto. %v", productResByte[0])
		return
	}

	return
}

//TODO: implementar taxa media para tipo de produto
//TODO: produto resume armazenar no shop
//TODO: essa classe calcula tudo
//TODO: classe promotions que  inclui a data de validade
