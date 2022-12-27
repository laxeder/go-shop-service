package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// restaura um endereço com status deletado
func RestoreAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	uid := ctx.Params("uid")

	addressDatabase, err := address.Repository().GetUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar endereço. (%v)", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS032"))
	}

	// verifica o status do endereço
	if addressDatabase.Status != address.Disabled {
		log.Error().Msgf("Este endereço já está ativo no sistema. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS033", "Este endereço já está ativo no sistema."))
	}

	// muda o status do endereço para ativo
	addressDatabase.Uid = uid
	addressDatabase.Status = address.Enabled
	addressDatabase.UpdatedAt = date.NowUTC()

	// salva as alterações na base de dados
	err = address.Repository().Restore(addressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS034", "O formado dos dados envidados está incorreto."))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
