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
		log.Error().Err(err).Msgf("O formado dos dados envidados está incorreto. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "GSS003", "O formado dos dados envidados está incorreto."))
	}

	//TODO:validar campo account
	// valida os campos de entrada
	// checkAccount := accountBody.Valid()
	// if checkAccount.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkAccount)
	// }

	//!##################################################################################################################//
	//! VERIFICA SE O DOCUMENTO DA CONTA EXISTE NA BASE DE DADOS 														 //
	//!##################################################################################################################//
	accountData, err := account.Repository().GetByUid(accountBody.Uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar encontrar a conta %v no repositório", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS004"))
	}

	// verifica se a conta está desabilitada
	if accountData.Status == account.Disabled {
		log.Error().Msgf("Esta conta (%v) está desabilitada por tempo indeterminado.", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS005", "Esta conta está desabilitada por tempo indeterminado."))
	}

	// verifica se existe uma uuid válida
	if len(accountData.Uuid) > 0 {
		log.Error().Msgf("Este documento já existe na nossa base de dados. (%v)", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS006", "Este documento já existe na nossa base de dados."))
	}

	// verifica se o documento existe
	if len(accountData.Uuid) > 0 {
		log.Error().Msgf("Este documento (%v) já existe na nossa base de dados.", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS007", "Este documento já existe na nossa base de dados."))
	}

	//!##################################################################################################################//
	//! CRIA UMA NOVA CONTA DE USUÁRIO E AMAZENA NA BASE DE DADOS 														 //
	//!##################################################################################################################//
	//accountBody.NewUuid()
	accountBody.Status = account.Enabled
	accountBody.CreatedAt = date.NowUTC()
	accountBody.UpdatedAt = date.NowUTC()

	// armazena na base de dados
	err = account.Repository().Save(accountBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", accountBody.Uuid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS008"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
