package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/utils/date"
)

func CreateFreight(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	freightBody, err := freight.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS060", "O formado dos dados envidados está incorreto."))
	}

	freightData, err := freight.Repository().GetByUid(freightBody.Uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar encontrar o frete %v no repositório", freightBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS061"))
	}

	if freightData.Status == freight.Disabled {
		log.Error().Msgf("Este frete (%v) está desabilitada por tempo indeterminado.", freightBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS062", "Este frete está desabilitado por tempo indeterminado."))
	}

	// verifica se existe uma uuid válida
	if len(freightData.Uid) > 0 {
		log.Error().Msgf("Este uid já existe na nossa base de dados. (%v)", freightBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS063", "Este uid já existe na nossa base de dados."))
	}

	// verifica se o documento existe
	if len(freightData.Uid) > 0 {
		log.Error().Msgf("Este uid (%v) já existe na nossa base de dados.", freightBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS064", "Este uid já existe na nossa base de dados."))
	}

	freightBody.NewUid()
	freightBody.Status = freight.Enabled
	freightBody.CreatedAt = date.NowUTC()
	freightBody.UpdatedAt = date.NowUTC()

	err = freight.Repository().Save(freightBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do frete %v", freightBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS065"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
