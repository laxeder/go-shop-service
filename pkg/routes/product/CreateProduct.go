package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// cria um novo produtos na base de ddaos
func CreateProduct(ctx *fiber.Ctx) error {

	var log = logger.New()

	body := ctx.Body()

	// transforma o json em Struct
	productBody, err := product.New(body)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.Error(400, "BLC002", "Os campos enviados estão incorretos."))
	}

	// Cria um ID para o produto
	productBody.NewUid()

	//!##################################################################################################################//
	//! VERIFICA SE O UID DO PRODUTO EXISTE NA BASE DE DADOS
	//!##################################################################################################################//
	productDatabase, err := product.Repository().GetUid(productBody.Uid)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. %v", err)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC031"))
	}

	// verifica se o produto está desabilitado
	if productDatabase.Status == product.Disabled {
		log.Error().Msgf("Este produto (%v) está desabilitado por tempo indeterminado.", productBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "BLC032", "Este produto está desabilitado por tempo indeterminado."))
	}

	// verifica se o produto existe
	if len(productDatabase.Uid) > 0 {
		log.Error().Msgf("Este produto (%v) já existe na nossa base de dados.", productBody.Uid)
		return response.Ctx(ctx).Result(response.Error(400, "BLC034", "Este produto já existe na nossa base de dados."))
	}

	//!##################################################################################################################//
	//! CRIAR UM NOVO PRODUTO E ARMAZENA NA BASE DE DADOS
	//!##################################################################################################################//

	productBody.Status = product.Enabled
	productBody.CreatedAt = date.NowUTC()
	productBody.UpdatedAt = date.NowUTC()

	// armazena o produtos na base de dados
	err = product.Repository().Save(productBody)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar repositório do produtos %v", productBody.Uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC003"))
	}

	return response.Ctx(ctx).Result(response.Success(201))
}
