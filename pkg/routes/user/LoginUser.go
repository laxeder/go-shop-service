package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/jwt"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func LoginUser(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	userBody := &user.User{}

	err := utils.InjectBytes(body, userBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body no user. %s", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS120", "O formado dos dados envidados está incorreto."))
	}

	userData, err := user.Repository().GetByEmail(userBody.Email)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário pelo email (%v).", userBody.Email)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS121"))
	}

	if userData == nil {
		log.Error().Msgf("Usuário não encontrado. %v", userBody.Email)
		return response.Ctx(ctx).Result(response.Error(400, "GSS122", "Esse usuário não foi encontrado na base de dados."))
	}

	userDataPass, err := user.Repository().GetPassword(userData.Uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter usuário (%v).", userBody.Email)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS121"))
	}

	if userDataPass == nil {
		log.Error().Err(err).Msgf("Usuário não encontrado (%v).", userData.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS189", "Esse usuário não foi encontrado na base de dados."))
	}

	hashPassword := utils.NewHashPassword(userDataPass.Salt, userBody.Password)

	if userDataPass.Password != hashPassword {
		return response.Ctx(ctx).Result(response.Error(400, "GSS123", "O email ou a senha está incorreto."))
	}

	token := jwt.New(userData).Token

	result := fiber.Map{"user": userData, "token": token}

	return response.Ctx(ctx).Result(response.Success(200, result))

}
