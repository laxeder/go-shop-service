package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func UserStatusCkeck(ctx *fiber.Ctx) error {

	// var log = logger.New()

	// document := ctx.Params("document")

	// userData, err := user.Repository().GetDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("Não foi possível encontrar o usuário na base de dados. (%v)", document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("BLC057"))
	// }

	// if userData.Status == user.Disabled {
	// 	log.Error().Msgf("Esta conta está desabilitada por tempo indeterminado. (%v)", document)
	// 	return response.Ctx(ctx).Result(response.Error(404, "BLC058", "Esta conta está desabilitada por tempo indeterminado."))
	// }

	return ctx.Next()
}
