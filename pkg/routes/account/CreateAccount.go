package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func CreateAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	accountBody := &account.Account{}

	err := utils.InjectBytes(body, accountBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body na conta.")
		return response.Ctx(ctx).Result(response.Error(400, "GSS003", "O formado dos dados envidados está incorreto."))
	}

	accountData, err := account.Repository().Get(accountBody.Uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter a conta (%v)", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS004"))
	}

	if accountData != nil {
		log.Error().Msgf("Email já está registrado (%v).", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS005", "Essa conta já está registrada na base de dados."))
	}

	err = account.Repository().Save(accountBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar salvar conta (%v)", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS008"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
