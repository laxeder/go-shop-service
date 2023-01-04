package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func UpdateUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")
	body := ctx.Body()
	userBody := &user.User{}

	err := utils.InjectBytes(body, userBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body no user. %s", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS130", "O formado dos dados envidados está incorreto."))
	}

	userData, err := user.Repository().Get(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário  (%v).", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS131"))
	}

	if userData == nil {
		log.Error().Err(err).Msgf("Usuário não encontrado (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS183", "Esse usuário não foi encontrado na base de dados."))
	}

	if userBody.Email != "" {
		userEmailData, err := user.Repository().GetByEmail(userBody.Email)

		if err != nil {
			log.Error().Err(err).Msgf("Erro ao tentar obter usuário pelo email (%v).", userBody.Email)
			return response.Ctx(ctx).Result(response.ErrorDefault("GSS187"))
		}

		if userEmailData != nil && userEmailData.Uuid != userData.Uuid {
			log.Error().Msgf("Email já está registrado (%v).", userBody.Email)
			return response.Ctx(ctx).Result(response.Error(400, "GSS188", "Esse email já está registrado na base de dados."))
		}
	}

	utils.Inject(userBody, userData)

	err = user.Repository().Update(userData)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar usuário %v", userBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS132"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
