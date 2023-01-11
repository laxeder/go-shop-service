package freight

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/utils/tokens"
	"github.com/rs/zerolog/log"
)

type Freight struct {
	Uid             string      `json:"uid,omitempty" redis:"uid,omitempty"`
	ZipcodeSender   string      `json:"zipcode_sender,omitempty" redis:"zipcode_sender,omitempty"`
	ZipcodeReceiver string      `json:"zipcode_receiver,omitempty" redis:"zipcode_receiver,omitempty"`
	Type            FreightType `json:"type,omitempty" redis:"type,omitempty"`
	Price           int         `json:"price,omitempty" redis:"price,omitempty"`
	Weight          int         `json:"weight,omitempty" redis:"weight,omitempty"`
	Heigth          int         `json:"heigth,omitempty" redis:"heigth,omitempty"`
	Width           int         `json:"width,omitempty" redis:"width,omitempty"`
	Lenght          int         `json:"lenght,omitempty" redis:"lenght,omitempty"`
	Status          Status      `json:"status,omitempty" redis:"status,omitempty"`
	UpdatedAt       string      `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
	CreatedAt       string      `json:"created_at,omitempty" redis:"created_at,omitempty"`
}

func New(freightByte ...[]byte) (freight *Freight, err error) {
	freight = &Freight{}
	err = nil

	if len(freightByte) == 0 {
		return
	}

	err = json.Unmarshal(freightByte[0], freight)
	if err != nil {
		log.Error().Err(err).Msgf("O json do frete está incorreto. %s", freightByte[0])
		return
	}

	return
}

func (f *Freight) Calc() int {
	f.Price = 1000

	return f.Price
}

func (f *Freight) NewUid() string {
	f.Uid = tokens.Nonce()
	return f.Uid
}

func (f *Freight) Inject(freight *Freight) *Freight {

	if freight.Uid != "" {
		f.Uid = freight.Uid
	}

	if freight.ZipcodeSender != "" {
		f.ZipcodeSender = freight.ZipcodeSender
	}

	if freight.ZipcodeReceiver != "" {
		f.ZipcodeReceiver = freight.ZipcodeReceiver
	}

	if freight.Type != "" {
		f.Type = freight.Type
	}

	if freight.Price != 0 {
		f.Price = freight.Price
	}

	if freight.Weight != 0 {
		f.Weight = freight.Weight
	}

	if freight.Heigth != 0 {
		f.Heigth = freight.Heigth
	}

	if freight.Weight != 0 {
		f.Weight = freight.Weight
	}

	if freight.Lenght != 0 {
		f.Lenght = freight.Lenght
	}

	if freight.CreatedAt != "" {
		f.CreatedAt = freight.CreatedAt
	}

	if freight.UpdatedAt != "" {
		f.UpdatedAt = freight.UpdatedAt
	}

	return f
}
