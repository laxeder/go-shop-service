package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

// cria um novo usuário na base de ddaos
func CreateUser(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	// transforma o json em Struct
	userBody, err := user.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos ou json está mal formatado. %s", userBody)
		return response.Ctx(ctx).Result(response.Error(400, "BLC002", "Os campos enviados estão incorretos ou json está mal formatado."))
	}

	// TODO validar
	// valida os campos de entrada
	// checkUser := userBody.Valid()
	// if checkUser.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkUser)
	// }

	//!##################################################################################################################//
	//! VERIFICA SE O DOCUMENTO DO USUÁRIO EXISTE NA BASE DE DADOS
	//!##################################################################################################################//
	userDatabase, err := user.Repository().GetDocument(userBody.Document)
	if err != nil {
		log.Error().Err(err).Msg("Os campos enviados estão incorretos.") // passar o documento
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC031"))
	}

	// verifica se a conta está desabilitada
	if userDatabase.Status == user.Disabled {
		log.Error().Msgf("Esta conta (%v) está desabilitada por tempo indeterminado.", userBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC032", "Esta conta está desabilitada por tempo indeterminado."))
	}

	// verifica se o documento existe
	if len(userDatabase.Document) > 0 {
		log.Error().Msgf("Este documento (%v) já existe na nossa base de dados.", userBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC034", "Este documento já existe na nossa base de dados."))
	}

	//!##################################################################################################################//
	//! CRIAR UM NOVO USUÁRIO E ARMAZENA NA BASE DE DADOS
	//!##################################################################################################################//

	// cria um novo usuário
	userBody.NewUuid()
	userBody.NewSalt()
	userBody.NewHashPassword()
	userBody.SetFullname()

	userBody.Status = user.Enabled
	userBody.CreatedAt = date.NowUTC()
	userBody.UpdatedAt = date.NowUTC()

	// armazena o usuário na base de dados
	err = user.Repository().Save(userBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", userBody.Document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC003"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
