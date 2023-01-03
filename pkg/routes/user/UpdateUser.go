package routes

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/shared"
)

func UpdateUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()
	uuid := ctx.Params("uuid")

	// converte json para struct
	userBody, err := user.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS130", "O formado dos dados envidados está incorreto."))
	}

	userDatabase, err := user.Repository().GetByUuid(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS131"))
	}

	spew.Dump(userDatabase)

	shared.Inject(userBody, userDatabase)

	spew.Dump(userDatabase)

	// injecta dos dados novos o lugar dos dsdos trazidos d abase de dados
	// userDatabase.Inject(userBody)
	// userDatabase.UpdatedAt = date.NowUTC()
	// userDatabase.SetFullname()

	// // guarda as alterações do usuário na base de dados
	// err = user.Repository().Update(userDatabase)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("Erro a tentar atualizar o repositório do usuário (%v)", userBody.Document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS132"))
	// }

	return response.Ctx(ctx).Result(response.Success(204))
}
