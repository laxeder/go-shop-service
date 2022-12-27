package freight

import (
	"encoding/json"

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

func New(userByte ...[]byte) (freight *Freight, err error) {
	freight = &Freight{}
	err = nil

	if len(userByte) == 0 {
		return
	}

	err = json.Unmarshal(userByte[0], freight)
	if err != nil {
		log.Error().Err(err).Msgf("O json do frete está incorreto. %s", userByte[0])
		return
	}

	return
}
