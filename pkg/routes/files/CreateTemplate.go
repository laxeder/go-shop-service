package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func CreateTemplate(ctx *fiber.Ctx) error {

	// 	var log = logger.New()

	// 	data := map[string]interface{}{
	// 		"name":      "igor",
	// 		"last_name": "barros",
	// 	}

	// 	pdf, err := shared.GeneratePDF(data, "exemple.tpl")
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("Erro ao tentar gerar um PDF")
	// 		return response.Ctx(ctx).Result(response.Error(400, "XXXX", "Ocorreu algum erro ao gerar um PDF"))
	// 	}

	// 	defer os.Remove(pdf)

	// 	return ctx.Download(pdf)
	return response.Ctx(ctx).Result(response.Success(201))

}
