package icp

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type ICP struct {
	Uuid             string `json:"uuid,omitempty" redis:"uuid,omitempty"`                           //  uuid do certificado
	Document         string `json:"document,omitempty" redis:"document,omitempty"`                   // documento do titular
	KeyPublic        string `json:"key_public,omitempty" redis:"key,omitempty"`                      // chave pública do titular
	Name             string `json:"name,omitempty" redis:"name,omitempty"`                           // naome do titular
	Email            string `json:"email,omitempty" redis:"email,omitempty"`                         // email do titular
	Validate         string `json:"validate,omitempty" redis:"validate,omitempty"`                   // validata do certificado
	ACName           string `json:"ac_name,omitempty" redis:"ac_name,omitempty"`                     // autoridade certificadora
	SerialNumber     string `json:"serial_number,omitempty" redis:"serial_number,omitempty"`         // numero de se'rie do certificado
	DigitalSignature string `json:"digital_signature,omitempty" redis:"digital_signature,omitempty"` // assinatura digital da autoridade cetificadora
	Status           Status `json:"status,omitempty" redis:"status,omitempty"`                       // status do certificado
	CreatedAt        string `json:"created_at,omitempty" redis:"created_at,omitempty"`               //
	UpdatedAt        string `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func New(icpByte ...[]byte) (icp *ICP, err error) {
	var log = logger.New()

	icp = &ICP{}
	err = nil

	if len(icpByte) == 0 {
		return
	}

	err = json.Unmarshal(icpByte[0], icp)
	if err != nil {
		log.Error().Err(err).Msgf("O json do certificado está incorreto. %v", icpByte[0])
		return
	}

	return
}

func (i *ICP) NewUuid() string {
	i.Uuid = uuid.New().String()
	return i.Uuid
}

func (i *ICP) Inject(icp *ICP) *ICP {

	if icp.Uuid != "" {
		i.Uuid = icp.Uuid
	}

	if icp.Document != "" {
		i.Document = icp.Document
	}

	if icp.KeyPublic != "" {
		i.KeyPublic = icp.KeyPublic
	}

	if icp.Name != "" {
		i.Name = icp.Name
	}

	if icp.Email != "" {
		i.Email = icp.Email
	}

	if icp.Validate != "" {
		i.Validate = icp.Validate
	}

	if icp.ACName != "" {
		i.ACName = icp.ACName
	}

	if icp.SerialNumber != "" {
		i.SerialNumber = icp.SerialNumber
	}

	if icp.Status != "" {
		i.Status = icp.Status
	}

	if icp.DigitalSignature != "" {
		i.DigitalSignature = icp.DigitalSignature
	}
	if icp.CreatedAt != "" {
		i.CreatedAt = icp.CreatedAt
	}

	if icp.UpdatedAt != "" {
		i.UpdatedAt = icp.UpdatedAt
	}

	return i

}
