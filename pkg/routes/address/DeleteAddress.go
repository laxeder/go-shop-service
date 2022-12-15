package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// muda o status do usuário na base de dados
func DeleteAddress(ctx *fiber.Ctx) error {
	var log = logger.New()

	uid := ctx.Params("uid")

	addressDatabase, err := address.Repository().GetUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS087"))
	}

	// verifica o status do endereço
	if addressDatabase.Status != address.Enabled {
		log.Error().Msgf("Este endereço já está desativado no sistema. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS060", "Este endereço já está desativado no sistema."))
	}

	addressDatabase.Uid = uid
	addressDatabase.Status = address.Disabled
	addressDatabase.UpdatedAt = date.NowUTC()

	err = address.Repository().Delete(addressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto %v.", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
