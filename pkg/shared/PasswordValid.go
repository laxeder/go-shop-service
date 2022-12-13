package shared

import (
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/regex"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func PasswordValid(password, document string) *response.Result {

	var log = logger.New()

	if password == "" {
		log.Error().Msgf("O campo da senha não pode ser vazio. (%v)", document)
		return response.Error(400, "BLC023", "O campo da senha não pode ser vazio.")
	}

	if len(password) < 8 {
		log.Error().Msgf("O campo da senha precisa ter no mínimo 8 caracteres. (%v)", document)
		return response.Error(400, "BLC024", "O campo da senha precisa ter no mínimo 8 caracteres.")
	}

	if !regex.HasUpppercase.MatchString(password) {
		log.Error().Msgf("A senha precisa ter pelo menos um caracter maiúsculo. (%v)", document)
		return response.Error(400, "BLC025", "A senha precisa ter pelo menos um caracter maiúsculo.")
	}

	if !regex.HasLowercase.MatchString(password) {
		log.Error().Msgf("A senha precisa ter pelo menos um caracter minúsculo. (%v)", document)
		return response.Error(400, "BLC026", "A senha precisa ter pelo menos um caracter minúsculo.")
	}

	if !regex.HasNumber.MatchString(password) {
		log.Error().Msgf("A senha precisa ter pelo menos um número. (%v)", document)
		return response.Error(400, "BLC027", "A senha precisa ter pelo menos um número.")
	}

	if !regex.HasCharSpecials.MatchString(password) {
		log.Error().Msgf("A senha precisa ter pelo menos um caractere especial. (%v)", document)
		return response.Error(400, "BLC028", "A senha precisa ter pelo menos um caractere especial.")
	}

	return response.Success(200)
}
