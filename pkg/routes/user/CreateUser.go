package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

// cria um novo usuário na base de ddaos
func CreateUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	// transforma o json em Struct
	userBody, err := user.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos ou json está mal formatado. %s", userBody)
		return response.Ctx(ctx).Result(response.Error(400, "GSS111", "Os campos enviados estão incorretos ou json está mal formatado."))
	}

	userDatabase, err := user.Repository().GetUuid(userBody.Uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos (%v). %v", userBody.Document, err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS112"))
	}

	// verifica se a conta está desabilitada
	if userDatabase.Status == user.Disabled {
		log.Error().Msgf("Esta conta (%v) está desabilitada por tempo indeterminado.", userBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "GSS113", "Esta conta está desabilitada por tempo indeterminado."))
	}

	// verifica se o documento existe
	if len(userDatabase.Document) > 0 {
		log.Error().Msgf("Este documento (%v) já existe na nossa base de dados.", userBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "GSS114", "Este documento já existe na nossa base de dados."))
	}

	userBody.NewUuid()
	userBody.NewSalt()
	userBody.NewHashPassword()
	userBody.SetFullname()

	userBody.Status = user.Enabled
	userBody.CreatedAt = date.NowUTC()
	userBody.UpdatedAt = date.NowUTC()

	// armazena o usuário na base de dados
	err = user.Repository().Save(userBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", userBody.Document)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS115"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
