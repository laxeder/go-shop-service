package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

// muda o status do usuário na base de dados
func DeleteUser(ctx *fiber.Ctx) error {
	var log = logger.New()

	document := ctx.Params("document")

	userDatabase, err := user.Repository().GetDocument(document)
	if err != nil {
		log.Error().Err(err).Msg("Os campos enviados estão incorretos.")
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC087"))
	}

	userDatabase.Status = user.Disabled
	userDatabase.UpdatedAt = date.NowUTC()

	err = user.Repository().Delete(userDatabase)
	if err != nil {
		log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
