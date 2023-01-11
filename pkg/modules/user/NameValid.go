package user

import (
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils/regex"
)

func NameValid(name string) *response.Result {

	var log = logger.New()

	if name == "" {
		log.Error().Msgf("O campo do nome não pode ser vazio. (%v)", name)
		return response.Error(400, "GSS143", "O campo do nome não pode ser vazio.")
	}

	if len(name) <= 2 {
		log.Error().Msgf("O nome não poder ser menor que 2 caracteres. (%v)", name)
		return response.Error(400, "GSS144", "O nome não poder ser menor que 2 caracteres.")
	}

	if regex.HasNumber.MatchString(name) {
		log.Error().Msgf("O nome não pode conter número. (%v)", name)
		return response.Error(400, "GSS145", "O nome não pode conter número.")
	}

	return response.Success(200)
}
