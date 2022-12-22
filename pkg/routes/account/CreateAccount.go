package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/account"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
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
	accountData, err := account.Repository().GetByUid(accountBody.Uid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar encontrar a conta %v no repositório", accountBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS004"))
	}

	// verifica se a conta está desabilitada
	if accountData.Status == account.Disabled {
		log.Error().Msgf("Esta conta (%v) está desabilitada por tempo indeterminado.", accountBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS005", "Esta conta está desabilitada por tempo indeterminado."))
	}

	// verifica se existe uma uuid válida
	if len(accountData.Uuid) > 0 {
		log.Error().Msgf("Este documento já existe na nossa base de dados. (%v)", accountBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS006", "Este documento já existe na nossa base de dados."))
	}

	// verifica se o documento existe
	if len(accountData.Uuid) > 0 {
		log.Error().Msgf("Este documento (%v) já existe na nossa base de dados.", accountBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "GSS007", "Este documento já existe na nossa base de dados."))
	}

	//!##################################################################################################################//
	//! CRIA UMA NOVA CONTA DE USUÁRIO E AMAZENA NA BASE DE DADOS 														 //
	//!##################################################################################################################//
	accountBody.NewUid()
	accountBody.Status = account.Enabled
	accountBody.CreatedAt = date.NowUTC()
	accountBody.UpdatedAt = date.NowUTC()

	// armazena na base de dados
	err = account.Repository().Save(accountBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do usuário %v", accountBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS008"))
	}

	// carrega o usuário da base de dados para atualizar as contas
	userDatabase, err := user.Repository().GetByUuid(accountBody.Uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar usuário (%v), (%v).", accountBody.Uid, err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS009"))
	}

	userDatabase.Accounts = append(userDatabase.Accounts, *accountBody)
	userDatabase.UpdatedAt = date.NowUTC()

	err = user.Repository().Update(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar contas do usuário (%v), (%v).", accountBody.Uid, err)
		return response.Ctx(ctx).Result(response.ErrorDefault("GSS010"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
