package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func CreateUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	userBody := &user.User{}

	err := utils.InjectBytes(body, userBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body no user (%s).", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS111", "O formado dos dados envidados está incorreto."))
	}

	userData, err := user.Repository().GetByEmail(userBody.Email)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário pelo email (%v).", userBody.Email)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS112"))
	}

	if userData != nil {
		log.Error().Msgf("Email já está registrado (%v).", userBody.Email)
		return response.Ctx(ctx).Result(response.Error(400, "GSS182", "Esse email já está registrado na base de dados."))
	}

	err = user.Repository().Save(userBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar salvar usuário (%v)", userBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS115"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
