package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func UpdateAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	uid := ctx.Params("uid")

	// converte json para struct
	addressBody, err := address.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS085", "O formado dos dados envidados está incorreto."))
	}

	// carrega o endereço da base de dados
	addressDatabase, err := address.Repository().GetByUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar endereço %v.", addressBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS081", "Erro ao tentar validar endereço."))
	}

	// injecta dos dados novos o lugar dos dsdos trazidos d abase de dados
	addressDatabase.Inject(addressBody)
	addressDatabase.UpdatedAt = date.NowUTC()

	// guarda as alterações do endereço na base de dados
	err = address.Repository().Update(addressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do endereço (%v)", addressBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS084"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
