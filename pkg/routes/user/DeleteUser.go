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
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC087"))
	}

	// verifica o status do usuário
	if userDatabase.Status != user.Enabled {
		log.Error().Msgf("Este usuário já está desativado no sistema. (%v)", document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC060", "Esta usuário já está desativado no sistema."))
	}

	userDatabase.Status = user.Disabled
	userDatabase.UpdatedAt = date.NowUTC()
	userDatabase.Document = document

	err = user.Repository().Delete(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC090"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
