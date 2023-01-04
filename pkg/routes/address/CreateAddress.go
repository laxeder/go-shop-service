package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/address"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// cria um novo endereço na base de ddaos
func CreateAddress(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	// transforma o json em Struct
	addressBody, err := address.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS022", "Os campos enviados estão incorretos."))
	}

	addressDatabase, err := address.Repository().GetUid(addressBody.Uid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS023"))
	}

	// verifica se a conta está desabilitada
	if addressDatabase.Status == address.Disabled {
		log.Error().Msgf("Esta conta (%v) está desabilitada por tempo indeterminado.", addressBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS024", "Esta conta está desabilitada por tempo indeterminado."))
	}

	// verifica se o documento existe
	if len(addressDatabase.Uid) > 0 {
		log.Error().Msgf("Este documento (%v) já existe na nossa base de dados.", addressBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS025", "Este documento já existe na nossa base de dados."))
	}

	addressBody.NewUid()

	addressBody.Status = address.Enabled
	addressBody.CreatedAt = date.NowUTC()
	addressBody.UpdatedAt = date.NowUTC()

	// armazena o endereço na base de dados
	err = address.Repository().Save(addressBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do endereço %v", addressBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS026"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
