package shared

import (
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/regex"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func NameValid(name string) *response.Result {

	var log = logger.New()

	if name == "" {
		log.Error().Msgf("O campo do nome não pode ser vazio. (%v)", name)
		return response.Error(400, "GSS004", "O campo do nome não pode ser vazio.")
	}

	if len(name) <= 2 {
		log.Error().Msgf("O nome não poder ser menor que 2 caracteres. (%v)", name)
		return response.Error(400, "GSS005", "O nome não poder ser menor que 2 caracteres.")
	}

	if regex.HasNumber.MatchString(name) {
		log.Error().Msgf("O nome não pode conter número. (%v)", name)
		return response.Error(400, "GSS006", "O nome não pode conter número.")
	}

	return response.Success(200)
}
