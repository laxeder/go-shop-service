package icp

import (
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/shared"
)

func (i *ICP) Valid() *response.Result {

	checkDocument := i.DocumentValid()
	if checkDocument.Status != 200 {
		return checkDocument
	}

	checkKeyPublic := i.KeyPublicValid()
	if checkKeyPublic.Status != 200 {
		return checkKeyPublic
	}

	checkName := i.NameValid()
	if checkName.Status != 200 {
		return checkName
	}

	checkEmail := i.EmailValid()
	if checkEmail.Status != 200 {
		return checkEmail
	}

	checkValidate := i.ValidateValid()
	if checkValidate.Status != 200 {
		return checkValidate
	}

	checkACName := i.ACNameValid()
	if checkACName.Status != 200 {
		return checkACName
	}

	checkSerialNumber := i.SerialNumberValid()
	if checkSerialNumber.Status != 200 {
		return checkSerialNumber
	}

	checkDigitalSignature := i.DigitalSignatureValid()
	if checkDigitalSignature.Status != 200 {
		return checkDigitalSignature
	}

	return response.Success(200)
}

func (i *ICP) DocumentValid() *response.Result {
	return shared.DocumentValid(i.Document)
}

func (i *ICP) KeyPublicValid() *response.Result {

	var log = logger.New()

	if len(i.KeyPublic) < 64 {
		log.Error().Msgf("A chave pública precisa ter no mínomo 64 caracteres. (%v)", i.Document)
		return response.Error(400, "BLC224", "A chave pública precisa ter no mínomo 64 caracteres.")
	}

	return response.Success(200)
}

func (i *ICP) NameValid() *response.Result {
	return shared.NameValid(i.Name)
}

func (i *ICP) EmailValid() *response.Result {
	return shared.EmailValid(i.Email)
}

func (i *ICP) ValidateValid() *response.Result {
	var log = logger.New()

	if !date.BRValid(i.Validate) {
		log.Error().Msgf("Data de validate do certificado está inválida: (%v)", i.Document)
		return response.Error(400, "BLC230", "Data de validate do certificado está inválida.")
	}

	return response.Success(200)
}

func (i *ICP) ACNameValid() *response.Result {
	return shared.NameValid(i.ACName)
}

func (i *ICP) SerialNumberValid() *response.Result {

	var log = logger.New()

	if i.SerialNumber == "" {
		log.Error().Msgf("O numero de serie do certificado não pode ser vazio. (%v)", i.Document)
		return response.Error(400, "BLC234", "O numero de serie do certificadonão pode ser vazio.")
	}

	if len(i.SerialNumber) <= 18 {
		log.Error().Msgf("O numero de serie do certificado não poder ser menor que 18 caracteres. (%v)", i.Document)
		return response.Error(400, "BLC235", "O numero de serie do certificado não poder ser menor que 18 caracteres.")
	}

	return response.Success(200)
}

func (i *ICP) DigitalSignatureValid() *response.Result {

	var log = logger.New()

	if i.DigitalSignature == "" {
		log.Error().Msgf("A assinatura digital da autoridade certificadora não pode ser vazia. (%v)", i.Document)
		return response.Error(400, "BLC236", "A assinatura digital da autoridade certificadora não pode ser vazia.")
	}

	if len(i.DigitalSignature) <= 18 {
		log.Error().Msgf("A assinatura digital da autoridade certificadorao não poder ser menor que 18 caracteres. (%v)", i.Document)
		return response.Error(400, "BLC237", "A assinatura digital da autoridade certificadora não poder ser menor que 18 caracteres.")
	}

	return response.Success(200)
}
