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
	uid := ctx.Params("uid")

	// converte json para struct
	accountBody, err := account.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS019", "O formado dos dados envidados está incorreto."))
	}

	// valida os campos enviados
	// checkAccount := accountBody.Valid()
	// if checkAccount.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkAccount)
	// }

	// carrega o usuário da base de dados

	accountDatabase, err := account.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar usuário %v.", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS020", "Erro ao tentar validar usuário."))
	}

	// formata a atualização
	accountDatabase.Inject(accountBody)
	accountDatabase.Birthday = date.BRToUTC(accountBody.Birthday)
	accountDatabase.UpdatedAt = date.NowUTC()

	// guarda as alterações na base de dados
	err = account.Repository().Update(accountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar encontrar o usuário %v no repositório", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS021"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
