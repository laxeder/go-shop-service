package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func ListICPsByDocument(ctx *fiber.Ctx) error {

	// var log = logger.New()

	// icpsData, err := icp.Repository().GetList()
	// if err != nil {
	// 	log.Error().Err(err).Msg("Erro ao tentar listar certificados.")
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("BLC246"))
	// }

	// return response.Ctx(ctx).Result(response.Success(200, icpsData))
	return response.Ctx(ctx).Result(response.Success(200))

}
