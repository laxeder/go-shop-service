package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// cria uma nova conta com base num usuário
func CreateAccount(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	// transforma o json em  struct
	accountBody, err := account.New(body)
	if err != nil {
		log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
		return response.Ctx(ctx).Result(response.Error(400, "BLC104", "O formado dos dados envidados está incorreto."))
	}

	//TODO:validar campo account
	// valida os campos de entrada
	// checkAccount := accountBody.Valid()
	// if checkAccount.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkAccount)
	// }

	//!##################################################################################################################//
	//! VERIFICA SE O DOCUMENTO DO USUÁRIO EXISTE NA BASE DE DADOS 														 //
	//!##################################################################################################################//
	accountData, err := account.Repository().GetDocument(accountBody.Document)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar encontrar o usuário %v no repositório", accountBody.Document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC105"))
	}

	// verifica se a conta está desabilitada
	if accountData.Status == account.Disabled {
		log.Error().Msgf("Esta conta (%v) está desabilitada por tempo indeterminado.", accountBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC032", "Esta conta está desabilitada por tempo indeterminado."))
	}

	// verifica se existe uma uuid válida
	if len(accountData.Uuid) > 0 {
		log.Error().Msgf("Este documento já existe na nossa base de dados. (%v)", accountBody.Document)
		return response.Ctx(ctx).Result(response.Error(400, "BLC106", "Este documento já existe na nossa base de dados."))
	}

	//!##################################################################################################################//
	//! CRIA UMA NOVA CONTA DE USUÁRIO E AMAZENA NA BASE DE DADOS 														 //
	//!##################################################################################################################//
	accountBody.NewUuid()
	accountBody.CreatedAt = date.NowUTC()
	accountBody.UpdatedAt = date.NowUTC()

	// armazena na base de dados
	err = account.Repository().Save(accountBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", accountBody.Document)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC107"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
