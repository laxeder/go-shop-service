package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// atualiza dados da conta
func UpdateAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	document := ctx.Params("document")

	// converte json para struct
	accountBody, err := account.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC164", "O formado dos dados envidados está incorreto."))
	}

	// valida os campos enviados
	// checkAccount := accountBody.Valid()
	// if checkAccount.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkAccount)
	// }

	// carrega o usuário da base de dados
	accountData, err := account.Repository().GetByDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar usuário %v.", accountBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC081", "Erro ao tentar validar usuário."))
	}

	// formata a atualização
	accountData.Inject(accountBody)
	accountData.Birthday = date.BRToUTC(accountBody.Birthday)
	accountData.UpdatedAt = date.NowUTC()

	// guarda as alterações na base de dados
	err = account.Repository().Update(accountData)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar encontrar o usuário %v no repositório", document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC167"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
