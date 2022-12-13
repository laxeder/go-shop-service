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
	document := ctx.Params("document")

	// iniicia a struct do usuário com password
	userPassowrd := &user.UserPassword{}

	// converte o json para struct
	err := json.Unmarshal(body, userPassowrd)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "O formado dos dados envidados está incorreto."))
	}

	// monta a struct de usuário
	userBody := &user.User{Password: userPassowrd.NewPassword, ConfirmPassword: userPassowrd.NewConfirmPassword}

	// compara o password
	if userPassowrd.Password == userPassowrd.NewPassword {
		log.Error().Msgf("A nova senha não pode ser idêntica a senha atual.(%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC077", "A nova senha não pode ser idêntica a senha atual."))
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

	// carrega os dados do usuário da base de dados
	userDatabase, err := user.Repository().GetPasswordByDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar carregar usuário %v.", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC093"))
	}

	// injeta os novos valores no lugar dos dados recuperados da base de dados
	userDatabase.Inject(userBody)
	userDatabase.NewSalt()
	userDatabase.NewHashPassword()
	userDatabase.UpdatedAt = date.NowUTC()

	// guarda a alterações na base de dados do usuário
	err = user.Repository().SavePassowrd(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o repositório do usuário %v", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC096"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
