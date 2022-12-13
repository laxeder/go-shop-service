package shared

import (
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/regex"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

func DocumentValid(document string) *response.Result {

	var log = logger.New()

	if document == "" {
		log.Error().Msgf("O campo do documento não pode ser vazio. (%v)", document)
		return response.Error(400, "BLC010", "O campo do documento não pode ser vazio.")
	}

	if regex.HasLetter.MatchString(document) {
		log.Error().Msgf("O documento não pode conter letras. (%v)", document)
		return response.Error(400, "BLC011", "O documento não pode conter letras.")
	}

	doc := str.DocumentClean(document)

	if len(doc) < 11 {
		log.Error().Msgf("O documento precisa ter no mínimo 11 números. (%v)", document)
		return response.Error(400, "BLC012", fmt.Sprintf("O documento %v precisa ter no mínimo 11 números.", document))
	}

	if len(doc) > 14 {
		log.Error().Msgf("O documento não pode ter mais que  14 números. (%v)", document)
		return response.Error(400, "BLC013", "O documento não pode ter mais que 14 números.")
	}

	if !str.DocumentValid(doc) {
		log.Error().Msgf("Documento inválido! Favor digitar um documento válido. (%v)", document)
		return response.Error(400, "BLC014", "Documento inválido! Favor digitar um documento válido.")
	}

	return response.Success(200)
}
