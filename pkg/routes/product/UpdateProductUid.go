package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func UpdateProductUid(ctx *fiber.Ctx) error {
	var log = logger.New()

	body := ctx.Body()
	uid := ctx.Params("uid")

	// inicia a struct do produto
	productUid := &product.ProductUid{}

	// converte o json para struct
	err := json.Unmarshal(body, productUid)
	if err != nil {
		log.Error().Err(err).Msgf("O formado dos dados enviados está incorreto. (%v)", uid)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "O formado dos dados enviados está incorreto."))
	}

	// carrega os dados do produto da base de dados
	productDatabase, err := product.Repository().GetUid(uid)
	if err != nil {
		log.Error().Err(err).Msgf("Produto não encontrado %v.", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC093"))
	}

	if productUid.OldUid != productDatabase.Uid {
		log.Error().Err(err).Msgf("Uid antigo errado %v.", uid)
		return response.Ctx(ctx).Result(response.Error(400, "BLC091", "uid antigo está incorreto."))
	}

	// define o novo uid
	productDatabase.SetUid(productUid.Uid)
	productDatabase.UpdatedAt = date.NowUTC()

	// guarda a alterações na base de dados do produto
	err = product.Repository().SaveUid(uid, productDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar o repositório do produto %v", uid)
		return response.Ctx(ctx).Result(response.ErrorDefault("BLC096"))
	}

	return response.Ctx(ctx).Result(response.Success(204))
}
