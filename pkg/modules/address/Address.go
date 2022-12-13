package address

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

type Address struct {
	Uuid         string `json:"uuid,omitempty" redis:"uuid,omitempty"`
	Document     string `json:"document,omitempty" redis:"document,omitempty"`
	Number       string `json:"number,omitempty" redis:"number,omitempty"`
	Zip          string `json:"zip,omitempty" redis:"zip,omitempty"`
	Street       string `json:"street,omitempty" redis:"street,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty" redis:"neighborhood,omitempty"`
	City         string `json:"city,omitempty" redis:"city,omitempty"`
	State        string `json:"state,omitempty" redis:"state,omitempty"`
	Status       Status `json:"status,omitempty" redis:"status,omitempty"`
	CreatedAt    string `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func New(addressByte ...[]byte) (address *Address, err error) {
	var log = logger.New()

	address = &Address{}
	err = nil

	if len(addressByte) == 0 {
		return address, err
	}

	err = json.Unmarshal(addressByte[0], address)
	if err != nil {
		log.Error().Err(err).Msgf("O json do address está incorreto. %v", addressByte[0])
		return address, err
	}

	return address, err
}

func (a *Address) SetNumber(number string) string {
	a.Number = number
	return a.Number
}

func (a *Address) SetZip(zip string) string {
	a.Zip = zip
	return a.Zip
}

func (a *Address) SetStreet(street string) string {
	a.Street = street
	return a.Street
}

func (a *Address) SetNeighborhood(neighborhood string) string {
	a.Neighborhood = neighborhood
	return a.Neighborhood
}

func (a *Address) SetCity(city string) string {
	a.City = city
	return a.City

}

func (a *Address) SetState(state string) string {
	a.State = state
	return a.State
}

func (a *Address) SetStatus(status Status) Status {
	a.Status = status
	return status
}

func (a *Address) SetDocument(document string) string {
	a.Document = str.DocumentPad(document)
	return a.Document
}

func (a *Address) NewUuid() string {
	return uuid.New().String()
}

func (a *Address) SetUuid(uuid string) string {
	a.Uuid = uuid
	return a.Uuid
}

func (a *Address) SetCreatedAt(createdAt string) string {
	a.CreatedAt = createdAt
	return a.CreatedAt
}

func (a *Address) SetUpdatedAt(updatedAt string) string {
	a.UpdatedAt = updatedAt
	return a.UpdatedAt
}

func (a *Address) ToString() (string, error) {
	var log = logger.New()

	addressJson, err := json.Marshal(a)
	if err != nil {
		log.Error().Err(err).Msgf("A struct do address está incorreta. (%v)", a.Document)
		return "", err
	}
	return string(addressJson), nil
}

func (a *Address) Inject(address *Address) *Address {

	if address.Uuid != "" {
		a.Uuid = address.Uuid
	}

	if address.Document != "" {
		a.Document = address.Document
	}

	if address.Number != "" {
		a.Number = address.Number
	}

	if address.Zip != "" {
		a.Zip = address.Zip
	}

	if address.Street != "" {
		a.Street = address.Street
	}

	if address.Neighborhood != "" {
		a.Neighborhood = address.Neighborhood
	}

	if address.City != "" {
		a.City = address.City
	}

	if address.State != "" {
		a.State = address.State
	}

	if address.CreatedAt != "" {
		a.CreatedAt = address.CreatedAt
	}

	if address.UpdatedAt != "" {
		a.UpdatedAt = address.UpdatedAt
	}

	return a
}
