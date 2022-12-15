package product

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
)

var redisClient *redis.Client

func Repository() *Product {
	return &Product{}
}

func MarshalBinary(str []string) (data []byte) {
	var log = logger.New()

	data, err := json.Marshal(str)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tranformar array em bytes %s", str)
	}

	return
}

func (a *Product) Save(product *Product) (err error) {
	var log = logger.New()

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return
	}

	key := fmt.Sprintf("products:%v", product.Uid)
	categories := MarshalBinary(product.Categories)
	pictures := MarshalBinary(product.Pictures)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uid", product.Uid)
		rdb.HSet(ctx, key, "name", product.Name)
		rdb.HSet(ctx, key, "description", product.Description)
		rdb.HSet(ctx, key, "pictures", pictures)
		rdb.HSet(ctx, key, "categories", categories)
		rdb.HSet(ctx, key, "price", product.Price)
		rdb.HSet(ctx, key, "promotion", product.Promotion)
		rdb.HSet(ctx, key, "code", product.Code)
		rdb.HSet(ctx, key, "weight", product.Weight)
		rdb.HSet(ctx, key, "color", product.Color)
		rdb.HSet(ctx, key, "updated_at", product.UpdatedAt)
		rdb.HSet(ctx, key, "created_at", product.CreatedAt)
		return err
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir o produto no redis. (%v)", product.Uid)
		return
	}

	return
}

func (a *Product) Update(product *Product) (err error) {
	var log = logger.New()

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return
	}

	key := fmt.Sprintf("products:%v", product.Uid)
	categories := MarshalBinary(product.Categories)
	pictures := MarshalBinary(product.Pictures)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uid", product.Uid)
		rdb.HSet(ctx, key, "name", product.Name)
		rdb.HSet(ctx, key, "description", product.Description)
		rdb.HSet(ctx, key, "pictures", pictures)
		rdb.HSet(ctx, key, "categories", categories)
		rdb.HSet(ctx, key, "price", product.Price)
		rdb.HSet(ctx, key, "promotion", product.Promotion)
		rdb.HSet(ctx, key, "code", product.Code)
		rdb.HSet(ctx, key, "weight", product.Weight)
		rdb.HSet(ctx, key, "color", product.Color)
		rdb.HSet(ctx, key, "updated_at", product.UpdatedAt)
		rdb.HSet(ctx, key, "created_at", product.CreatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o produto no redis. (%v)", product.Uid)
		return
	}

	return
}

func (u *Product) GetUid(uid string) (product *Product, err error) {

	var log = logger.New()

	product = &Product{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return nil, err
	}

	key := fmt.Sprintf("products:%v", uid)
	res := redisClient.HMGet(ctx, key, "uid", "status")
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o produto com o uid: %v.", uid)
		return u, err
	}

	err = res.Scan(product)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um produto válido com o uid %v.", uid)
		return nil, err
	}

	if product.Status == Disabled {
		product = &Product{Status: Disabled}
		return product, nil
	}

	return product, nil
}

func (a *Product) GetByUid(uid string) (product *Product, err error) {

	var log = logger.New()

	product = &Product{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return nil, err
	}

	key := fmt.Sprintf("products:%v", uid)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o produto com o uid: %v.", uid)
		return nil, err
	}

	err = res.Scan(product)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear o produto válido para o uid %v.", uid)
		return nil, err
	}

	if product.Status == Disabled {
		product = &Product{Status: Disabled}
		return product, nil
	}

	return product, nil
}

func (u *Product) SaveUid(uid string, product *Product) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return err
	}

	key := fmt.Sprintf("products:%v", uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uid", product.Uid)
		rdb.HSet(ctx, key, "name", product.Name)
		rdb.HSet(ctx, key, "description", product.Description)
		rdb.HSet(ctx, key, "pictures", product.Pictures)
		rdb.HSet(ctx, key, "categories", product.Categories)
		rdb.HSet(ctx, key, "price", product.Price)
		rdb.HSet(ctx, key, "promotion", product.Promotion)
		rdb.HSet(ctx, key, "code", product.Code)
		rdb.HSet(ctx, key, "weight", product.Weight)
		rdb.HSet(ctx, key, "color", product.Color)
		rdb.HSet(ctx, key, "updated_at", product.UpdatedAt)
		rdb.HSet(ctx, key, "created_at", product.CreatedAt)
		return err
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o produto com o uid %v no redis.", uid)
		return err
	}

	return nil
}

func (u *Product) Delete(product *Product) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return err
	}

	product.Status = Disabled

	key := fmt.Sprintf("products:%v", product.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(product.Status))
		rdb.HSet(ctx, key, "updated_at", product.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o produto com o uid %v no redis.", product.Uid)
		return err
	}

	return nil
}

func (u *Product) Restore(product *Product) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return err
	}

	product.Status = Enabled

	key := fmt.Sprintf("products:%v", product.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(product.Status))
		rdb.HSet(ctx, key, "updated_at", product.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível restaurar o produto com o uid %v no redis.", product.Uid)
		return err
	}

	return nil
}

func (u *Product) GetList() (products []Product, err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return nil, err
	}

	iter := redisClient.Scan(ctx, 0, "products:*", 0).Iterator()
	for iter.Next(ctx) {
		uid := strings.Replace(iter.Val(), "products:", "", 2)
		product, uErr := u.GetByUid(uid)

		if uErr != nil {
			continue
		}

		if product.Status == Disabled {
			continue
		}

		products = append(products, *product)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar aos produtos do banco de dados.")
		return nil, err
	}

	return products, nil
}
