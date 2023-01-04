package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func UpdateUserDocument(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	userBody := &user.UserDocument{}

	err := utils.InjectBytes(body, userBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body no user. %s", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS193", "O formado dos dados envidados está incorreto."))
	}

	userData, err := user.Repository().Get(userBody.Uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário  (%v).", userBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS133"))
	}

	if userData == nil {
		log.Error().Err(err).Msgf("Usuário não encontrado (%v).", userBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS190", "Esse usuário não foi encontrado na base de dados."))
	}

	userBody.GenerateDocument()

	if userBody.OldDocument != userData.Document {
		return response.Ctx(ctx).Result(response.Error(400, "GSS134", "O documento antigo está incorreto."))
	}

	userData.Document = userBody.NewDocument

	err = user.Repository().SaveDocument(userData)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar documento do usuário %v", userBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS135"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
