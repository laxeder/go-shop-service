package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func UserDocumentCheck(ctx *fiber.Ctx) error {

	// var log = logger.New()

	// document := ctx.Params("document")

	// if document == "" {
	// 	log.Error().Msg("O docuemnto do usuário não pode ser vazio.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS074", "O docuemnto do usuário não pode ser vazio."))
	// }

	// userBody := &user.User{Document: document}

	// checkDocument := userBody.DocumentValid()
	// if checkDocument.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkDocument)
	// }

	// userData, err := user.Repository().GetDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS075"))
	// }

	// if len(userData.Document) == 0 {
	// 	log.Error().Msg("O documento do usuário não foi encontrado.")
	// 	return response.Ctx(ctx).Result(response.Error(404, "GSS076", "O documento do usuário não foi encontrado."))
	// }

	return ctx.Next()
}
