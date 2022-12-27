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

func MarshalBinary(str []string) (data []byte) {
	var log = logger.New()

	data, err := json.Marshal(str)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tranformar array em bytes %s", str)
	}

	return
}

func UnmarshalBinary(bff []byte) []string {
	var log = logger.New()

	data := &[]string{}

	err := json.Unmarshal(bff, data)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tranformar bytes em array %s", bff)
	}

	return *data
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

	shopcart.ApplyFreightUid()
	shopcart.ApplyProductsUid()

	products := MarshalBinary(shopcart.ProductsUid)
	freights := MarshalBinary(shopcart.FreightsUid)

	key := fmt.Sprintf("shopcarts:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", uuid)
		rdb.HSet(ctx, key, "products_uid", products)
		rdb.HSet(ctx, key, "freights_uid", freights)
		rdb.HSet(ctx, key, "taxes", shopcart.Taxes)
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

	shopcart.ApplyFreightUid()
	shopcart.ApplyProductsUid()

	products := MarshalBinary(shopcart.ProductsUid)
	freights := MarshalBinary(shopcart.FreightsUid)

	key := fmt.Sprintf("shopcarts:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "products_uid", products)
		rdb.HSet(ctx, key, "freights_uid", freights)
		rdb.HSet(ctx, key, "taxes", shopcart.Taxes)
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
	shopcart.ProductsUid = UnmarshalBinary([]byte(res.Val()["products_uid"]))
	shopcart.FreightsUid = UnmarshalBinary([]byte(res.Val()["freights_uid"]))

	shopcart.ForEachProductsUid(func(uid string) {
		productDatabase, err := product.Repository().GetByUid(uid)

		if err != nil {
			log.Error().Err(err).Msgf("Erro ao buscar produto (%v) do carrinho de compras. %v", uid)
			return
		}

		if productDatabase.Uid == "" {
			log.Error().Msgf("Produto (%v) não existe. %v", uid)
			return
		}

		if productDatabase.Status == product.Disabled {
			productDatabase = &product.Product{Status: product.Disabled}
		}

		shopcart.Products = append(shopcart.Products, *productDatabase)
	})

	shopcart.ForEachFreightsUid(func(uid string) {
		freightDatabase, err := freight.Repository().GetByUid(uid)

		if err != nil {
			log.Error().Err(err).Msgf("Erro ao buscar frete (%v) do carrinho de compras. %v", uid)
			return
		}

		if freightDatabase.Uid == "" {
			log.Error().Msgf("Frete (%v) não encontrado. %v", uid, err)
			return
		}

		if freightDatabase.Status == freight.Disabled {
			freightDatabase = &freight.Freight{Status: freight.Disabled}
		}

		shopcart.Freights = append(shopcart.Freights, *freightDatabase)
	})

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
