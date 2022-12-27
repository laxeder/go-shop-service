package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/jwt"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func LoginUser(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()

	// transforma o json em Struct
	userBody, err := user.NewUserLogin(body)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos ou json está mal formatado. %s", userBody)
		return response.Ctx(ctx).Result(response.Error(400, "GSS120", "Os campos enviados estão incorretos ou json está mal formatado."))
	}

	userDatabase, err := user.Repository().GetByEmail(userBody.Email)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS121"))
	}

	if userDatabase == nil {
		log.Error().Msgf("Usuário não encontrado. %s", userBody)
		return response.Ctx(ctx).Result(response.Error(400, "GSS122", "Usuário não encontrado."))
	}

	hashPassword := userBody.HashPassword(userDatabase.Salt, userBody.Password)

	if userDatabase.Password != hashPassword {
		return response.Ctx(ctx).Result(response.Error(400, "GSS123", "Email ou senha incorretos."))
	}

	userDatabase.Password = ""
	userDatabase.ConfirmPassword = ""
	userDatabase.Salt = ""

	token := jwt.New(userDatabase).Token

	result := fiber.Map{"user": userDatabase, "token": token}

	return response.Ctx(ctx).Result(response.Success(200, result))

}
