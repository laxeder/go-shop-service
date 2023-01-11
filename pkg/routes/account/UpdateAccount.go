package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func UpdateAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")
	body := ctx.Body()
	accountBody := &account.Account{}

	err := utils.InjectBytes(body, accountBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body na conta. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS019", "O formado dos dados envidados está incorreto."))
	}

	accountData, err := account.Repository().Get(uuid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter conta %v.", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS020"))
	}

	if accountData == nil {
		log.Error().Err(err).Msgf("Conta não encontrada (%v).", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS183", "Essa conta não foi encontrada na base de dados."))
	}

	utils.Inject(accountBody, accountData)

	err = account.Repository().Update(accountData)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar conta (%v).", accountBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS021"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
