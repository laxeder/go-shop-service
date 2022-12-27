package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func UpdateUserDocument(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	uuid := ctx.Params("uuid")

	// iniicia a struct do usuário com password
	userDocument := &user.UserDocument{}

	// converte o json para struct
	err := json.Unmarshal(body, userDocument)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. (%v)", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS132", "O formado dos dados envidados está incorreto."))
	}

	userDatabase, err := user.Repository().GetByUuid(uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Documento não encontrado %v.", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS133"))
	}

	if userDocument.OldDocument != uuid {
		log.Error().Err(err).Msgf("Documento antigo errado %v.", uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS134", "Documento antigo está incorreto."))
	}

	// define o novo documento
	userDatabase.SetDocument(userDocument.Document)
	userDatabase.UpdatedAt = date.NowUTC()

	// guarda a alterações na base de dados do usuário
	err = user.Repository().SaveDocument(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o repositório do usuário %v", uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS135"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
