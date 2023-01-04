package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func UpdateUserPassword(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	userBody := &user.UserPassword{}

	err := utils.InjectBytes(body, userBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body no user. %s", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS136", "O formado dos dados envidados está incorreto."))
	}

	if userBody.NewPassword != userBody.ConfirmPassword {
		return response.Ctx(ctx).Result(response.Error(400, "GSS191", "A confirmação de senha está incorreta."))
	}

	userData, err := user.Repository().GetPassword(userBody.Uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário  (%v).", userBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS137"))
	}

	if userData == nil {
		log.Error().Err(err).Msgf("Usuário não encontrado (%v).", userBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS138", "Esse usuário não foi encontrado na base de dados."))
	}

	hashPassword := utils.NewHashPassword(userData.Salt, userBody.OldPassword)

	if userData.Password != hashPassword {
		return response.Ctx(ctx).Result(response.Error(400, "GSS139", "A senha antiga está incorreta."))
	}

	userData.Password = userBody.NewPassword

	err = user.Repository().SavePassword(userData)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar a senha do usuário %v", userBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS192"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
