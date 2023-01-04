package shopcart

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
)

var redisClient *redis.Client

func Repository() *ShopCart {
	return &ShopCart{}
}

func MarshalBinary(str []product.ProductResume) (data []byte) {
	var log = logger.New()

	data, err := json.Marshal(str)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tranformar array em bytes %s", str)
	}

	return
}

func UnmarshalBinary(bff []byte) (data any) {
	var log = logger.New()

	err := json.Unmarshal(bff, data)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tranformar bytes em array %s", bff)
	}

	return
}

func (s *ShopCart) Save(shopcart *ShopCart) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uuid := shopcart.Uuid

	redisClient, err = redisdb.New(redisdb.ShopCartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ShopCartDatabase)
		return err
	}

	shopcart.ProductsResume = shopcart.ReduceProducts(shopcart.Products)

	key := fmt.Sprintf("shopcarts:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", uuid)
		rdb.HSet(ctx, key, "products_resume", MarshalBinary(shopcart.ProductsResume))
		rdb.HSet(ctx, key, "taxes", shopcart.Taxes)
		rdb.HSet(ctx, key, "discount_percentage", shopcart.DiscountPercentage)
		rdb.HSet(ctx, key, "status", string(shopcart.Status))
		rdb.HSet(ctx, key, "last_acesses", shopcart.LastAcesses)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir o carrinho de compras com uuid %v no redis.", uuid)
		return err
	}

	return nil
}

func (s *ShopCart) Update(shopcart *ShopCart) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uuid := shopcart.Uuid

	redisClient, err = redisdb.New(redisdb.ShopCartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ShopCartDatabase)
		return err
	}

	shopcart.ProductsResume = shopcart.ReduceProducts(shopcart.Products)

	key := fmt.Sprintf("shopcarts:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "products_resume", MarshalBinary(shopcart.ProductsResume))
		rdb.HSet(ctx, key, "taxes", shopcart.Taxes)
		rdb.HSet(ctx, key, "discount_percentage", shopcart.DiscountPercentage)
		rdb.HSet(ctx, key, "last_acesses", shopcart.LastAcesses)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o carrinho de compras com uuid %v no redis.", uuid)
		return err
	}

	return nil
}

func (s *ShopCart) Delete(shopcart *ShopCart) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uuid := shopcart.Uuid

	redisClient, err = redisdb.New(redisdb.ShopCartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ShopCartDatabase)
		return err
	}

	shopcart.Status = Disabled

	key := fmt.Sprintf("shopcarts:%v", uuid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(shopcart.Status))
		rdb.HSet(ctx, key, "last_acesses", shopcart.LastAcesses)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o carrinho de compras com uuid %v no redis.", uuid)
		return err
	}

	return nil
}

func (s *ShopCart) Restore(shopcart *ShopCart) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uuid := shopcart.Uuid

	redisClient, err = redisdb.New(redisdb.ShopCartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ShopCartDatabase)
		return err
	}

	shopcart.Status = Enabled

	key := fmt.Sprintf("shopcarts:%v", uuid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(shopcart.Status))
		rdb.HSet(ctx, key, "last_acesses", shopcart.LastAcesses)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível retaurar o carrinho de compras com uuid %v no redis.", uuid)
		return err
	}

	return nil
}

func (s *ShopCart) GetUuid(uuid string) (shopcart *ShopCart, err error) {

	var log = logger.New()

	shopcart = &ShopCart{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.ShopCartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ShopCartDatabase)
		return nil, err
	}

	key := fmt.Sprintf("shopcarts:%v", uuid)
	res := redisClient.HMGet(ctx, key, "uid", "uuid", "status")

	err = res.Err()

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o carrinho de compras com uuid: %v.", uuid)
		return s, err
	}

	err = res.Scan(shopcart)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um carrinho de compras válido para o uuid %v.", uuid)
		return nil, err
	}

	if shopcart.Status == Disabled {
		shopcart = &ShopCart{Status: Disabled}
		return shopcart, nil
	}

	return shopcart, nil
}

func (s *ShopCart) GetByUuid(uuid string) (shopcart *ShopCart, err error) {

	var log = logger.New()

	shopcart = &ShopCart{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.ShopCartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ShopCartDatabase)
		return nil, err
	}

	key := fmt.Sprintf("shopcarts:%v", uuid)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o carrinho de compras com uuid: %v.", uuid)
		return nil, err
	}

	err = res.Scan(shopcart)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um carrinho de compras válido para o uuid %v.", uuid)
		return nil, err
	}

	if shopcart.Status == Disabled {
		shopcart = &ShopCart{Status: Disabled}
		return shopcart, nil
	}

	//? Convertendo bytes
	productRes, er := product.UnmarshalProductsResume([]byte(res.Val()["products_resume"]))

	if er != nil {
		return nil, er
	}

	shopcart.ProductsResume = *productRes

	var productsResume []product.ProductResume

	for _, productRes := range shopcart.ProductsResume {
		productDatabase, er := product.Repository().GetByUid(productRes.ProductUid)

		if er != nil {
			log.Error().Err(er).Msgf("Erro ao buscar produto (%v) do carrinho de compras.", productRes.ProductUid)
			continue
		}

		if productDatabase.Uid == "" {
			log.Error().Msgf("Produto (%v) não existe.", productRes.ProductUid)
			continue
		}

		if productDatabase.Status == product.Disabled {
			continue
		}

		freightDatabase, err := freight.Repository().GetByUid(productRes.FreightUid)

		if err != nil {
			log.Error().Err(err).Msgf("Erro ao buscar frete (%v) do carrinho de compras.", productRes.FreightUid)
			continue
		}

		if freightDatabase.Uid == "" {
			log.Error().Msgf("Frete (%v) não encontrado.", productRes.FreightUid)
			continue
		}

		if freightDatabase.Status == freight.Disabled {
			continue
		}

		productDatabase.Freight = *freightDatabase
		productDatabase.Freight.ZipcodeReceiver = productRes.ZipcodeReceiver
		productDatabase.Freight.Calc()

		productRes.Product = *productDatabase

		productsResume = append(productsResume, productRes)
	}

	shopcart.ProductsResume = productsResume

	shopcart.ExpandProducts(shopcart.ProductsResume)
	shopcart.FreightResume()
	shopcart.Resume()

	return shopcart, nil
}

func (s *ShopCart) GetList() (shopcarts []ShopCart, err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.ShopCartDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ShopCartDatabase)
		return nil, err
	}

	iter := redisClient.Scan(ctx, 0, "shopcarts:*", 0).Iterator()
	for iter.Next(ctx) {
		uuid := strings.Replace(iter.Val(), "shopcarts:", "", 2)
		shopcart, uErr := s.GetByUuid(uuid)

		if uErr != nil {
			continue
		}

		if shopcart.Status == Disabled {
			continue
		}

		shopcarts = append(shopcarts, *shopcart)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar os carrinhos de compras do banco de dados.")
		return nil, err
	}

	return shopcarts, nil
}
