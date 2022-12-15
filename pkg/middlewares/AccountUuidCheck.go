package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AccountUuidCheck(ctx *fiber.Ctx) error {
	// var log = logger.New()

	// document := ctx.Params("document")

	// accountDatabase, err := account.Repository().GetDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Erro ao tentar validar usuário.")
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS092"))
	// }

	// if len(accountDatabase.Uuid) == 0 {
	// 	log.Error().Err(err).Msg("Este documento não existe na nossa base de dados.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS094", "Este documento não existe na nossa base de dados."))
	// }

	return ctx.Next()
}
