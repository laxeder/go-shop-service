package category

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var redisClient *redis.Client
var log = logger.New()

func Repository() *Category {
	return &Category{}
}

func (c *Category) Save(category *Category) (err error) {

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.ProductDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.ProductDatabase)
		return
	}

	code := str.PadCategory(category.Code)

	key := fmt.Sprintf("categories:%v", code)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "name", category.Name)
		rdb.HSet(ctx, key, "description", category.Description)
		rdb.HSet(ctx, key, "code", category.Code)
		rdb.HSet(ctx, key, "status", string(category.Status))
		rdb.HSet(ctx, key, "updated_at", category.UpdatedAt)
		rdb.HSet(ctx, key, "created_at", category.CreatedAt)
		return err
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir a categoria no redis. (%v)", category.Code)
		return
	}

	return
}

func (c *Category) Update(category *Category) (err error) {

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.CategoriesDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.CategoriesDatabase)
		return
	}

	code := str.PadCategory(category.Code)

	key := fmt.Sprintf("categories:%v", code)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "name", category.Name)
		rdb.HSet(ctx, key, "description", category.Description)
		rdb.HSet(ctx, key, "updated_at", category.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar a categoria no redis. (%v)", category.Code)
		return
	}

	return
}

func (c *Category) GetByCode(code string) (category *Category, err error) {

	var log = logger.New()

	category = &Category{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.CategoriesDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.CategoriesDatabase)
		return nil, err
	}

	code = str.PadCategory(code)

	key := fmt.Sprintf("categories:%v", code)
	res := redisClient.HGetAll(ctx, key)

	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar a categoria com o code: %v.", code)
		return
	}

	err = res.Scan(category)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear a categoria com o code %v.", code)
		return
	}

	if category.Status == Disabled {
		category = &Category{Status: Disabled, Code: code}
		return
	}

	return
}

func (c *Category) Delete(category *Category) (err error) {

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.CategoriesDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.CategoriesDatabase)
		return err
	}

	code := str.PadCategory(category.Code)

	category.Status = Disabled

	key := fmt.Sprintf("categories:%v", code)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(category.Status))
		rdb.HSet(ctx, key, "updated_at", category.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar a categoria com o code %v no redis.", code)
		return err
	}

	return nil
}

func (c *Category) Restore(category *Category) (err error) {

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.CategoriesDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.CategoriesDatabase)
		return err
	}

	code := str.PadCategory(category.Code)

	category.Status = Enabled

	key := fmt.Sprintf("categories:%v", code)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(category.Status))
		rdb.HSet(ctx, key, "updated_at", category.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível restaurar a categoria com o code %v no redis.", code)
		return err
	}

	return nil
}

func (c *Category) GetList() (categories []Category, err error) {

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.CategoriesDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.CategoriesDatabase)
		return
	}

	iter := redisClient.Scan(ctx, 0, "categories:*", 0).Iterator()
	for iter.Next(ctx) {
		code := strings.Replace(iter.Val(), "categories:", "", 2)
		category, uErr := c.GetByCode(code)

		if uErr != nil {
			continue
		}

		if category.Status == Disabled {
			continue
		}

		categories = append(categories, *category)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar as categorias do banco de dados.")
		return
	}

	return
}

func (c *Category) GetKeys() (codes []string, err error) {

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.CategoriesDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.CategoriesDatabase)
		return
	}

	iter := redisClient.Scan(ctx, 0, "categories:*", 0).Iterator()
	for iter.Next(ctx) {
		codes = append(codes, strings.Replace(iter.Val(), "categories:", "", 2))
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar as categorias do banco de dados.")
		return
	}

	return
}
