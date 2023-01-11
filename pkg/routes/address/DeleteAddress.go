package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

func DeleteAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")
	uid := ctx.Params("uid")

	addressData, err := address.Repository().GetDataInfo(uuid, uid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter endereço (%v:%v).", uuid, uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS029"))
	}

	if addressData == nil {
		log.Error().Msgf("Endereço não encontrado (%v:%v).", uuid, uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS199", "Esse endereço não foi encontrado na base de dados."))
	}

	if addressData.Status != status.Enabled {
		log.Error().Msgf("Endereço já está desativado. (%v:%v)", uuid, uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS030", "Este endereço já está desativado na base de dados."))
	}

	err = address.Repository().Delete(uuid, uid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar deletar endereço (%v:%v).", uuid, uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS031"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
