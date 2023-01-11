package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func UpdateAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	uuid := ctx.Params("uuid")
	uid := ctx.Params("uid")
	body := ctx.Body()
	addressBody := &address.Address{}

	err := utils.InjectBytes(body, addressBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body no user (%s).", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS036", "O formado dos dados envidados está incorreto."))
	}

	addressData, err := address.Repository().Get(uuid, uid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter endereço (%v:%v).", uuid, uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS037"))
	}

	if addressData == nil {
		log.Error().Err(err).Msgf("Endereço não encontrado (%v:%v).", uuid, uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS221", "Esse endereço não foi encontrado na base de dados."))
	}

	utils.Inject(addressBody, addressData)

	err = address.Repository().Update(addressData)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o endereço (%v)", addressBody)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS038"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
