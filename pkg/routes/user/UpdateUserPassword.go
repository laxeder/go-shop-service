package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func UpdateUserPassword(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	uuid := ctx.Params("uuid")

	// iniicia a struct do usuário com password
	userPassowrd := &user.UserPassword{}

	// converte o json para struct
	err := json.Unmarshal(body, userPassowrd)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS091", "O formado dos dados envidados está incorreto."))
	}

	// monta a struct de usuário
	userBody := &user.User{Password: userPassowrd.NewPassword, ConfirmPassword: userPassowrd.NewConfirmPassword}

	// compara o password
	if userPassowrd.Password == userPassowrd.NewPassword {
		log.Error().Msgf("A nova senha não pode ser idêntica a senha atual.(%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS077", "A nova senha não pode ser idêntica a senha atual."))
	}

	// valida o password
	checkPassword := userBody.PasswordValid()
	if checkPassword.Status != 200 {
		return response.Ctx(ctx).Result(checkPassword)
	}

	// valida a confirmação
	checkConfirmPassword := userBody.ConfirmPasswordValid()
	if checkConfirmPassword.Status != 200 {
		return response.Ctx(ctx).Result(checkConfirmPassword)
	}

	userDatabase, err := user.Repository().GetByUuid(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS035"))
	}

	// injeta os novos valores no lugar dos dados recuperados da base de dados
	userDatabase.Inject(userBody)
	userDatabase.NewSalt()
	userDatabase.NewHashPassword()
	userDatabase.UpdatedAt = date.NowUTC()

	// guarda a alterações na base de dados do usuário
	err = user.Repository().SavePassowrd(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o repositório do usuário %v", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS096"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
