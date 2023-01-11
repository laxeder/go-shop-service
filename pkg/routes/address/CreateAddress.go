package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

func CreateAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	addressBody := &address.Address{}

	err := utils.InjectBytes(body, addressBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar injetar a body no endereço (%s).", body)
		return response.Ctx(ctx).Result(response.Error(400, "GSS022", "O formado dos dados envidados está incorreto."))
	}

	addressData, err := address.Repository().Get(addressBody.Uuid, addressBody.Uid)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar obter endereço (%v:%v).", addressBody.Uuid, addressBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS023"))
	}

	if addressData != nil {
		log.Error().Msgf("Endereço já está registrado (%v:%v).", addressBody.Uuid, addressBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS024", "Esse endereço já esta registrado na base de dados."))
	}

	err = address.Repository().Save(addressBody)

	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar salvar endereço (%v:%v).", addressBody.Uuid, addressBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS026"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
