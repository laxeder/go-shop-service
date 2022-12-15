package shared

import (
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/regex"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func EmailValid(email string) *response.Result {

	var log = logger.New()

	if email == "" {
		log.Error().Msgf("O campo do email não pode ser vazio. (%v)", email)
		return response.Error(400, "GSS015", "O campo do email não pode ser vazio.")
	}

	if len(email) <= 6 {
		log.Error().Msgf("O email não poder ser menor que 6 caracteres. (%v)", email)
		return response.Error(400, "GSS016", "O email não poder ser menor que 6 caracteres.")
	}

	if !regex.IsEmail.MatchString(email) {
		log.Error().Msgf("Email inválido, tente inserir outro email (%v)", email)
		return response.Error(400, "GSS017", "Email inválido, tente inserir outro email.")
	}

	return response.Success(200)
}
