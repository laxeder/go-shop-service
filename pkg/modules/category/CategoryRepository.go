package category

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

var redisClient *redis.Client

func Repository() *Category {
	return &Category{}
}

func (c *Category) Exists(code string, ignoreStatus bool) (bool, error) {

	categoryData, err := redisdb.GetDataInfo(redisdb.CategoryDatabase, fmt.Sprintf("categories:%v", code))

	if err != nil {
		return false, err
	}

	if categoryData == nil {
		return false, nil
	}

	if !ignoreStatus && categoryData.Status != status.Enabled {
		return false, nil
	}

	return true, nil
}

func (c *Category) Save(category *Category) (err error) {

	redisClient, err = redisdb.New(redisdb.CategoryDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	category.GenerateCode()

	key := fmt.Sprintf("categories:%v", category.Code)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "code", category.Code)
		rdb.HSet(ctx, key, "name", category.Name)
		rdb.HSet(ctx, key, "description", category.Description)

		redisdb.CreateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (c *Category) Update(category *Category) (err error) {

	exists, err := c.Exists(category.Code, false)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Category does not exist")
	}

	redisClient, err = redisdb.New(redisdb.CategoryDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("categories:%v", category.Code)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "name", category.Name)
		rdb.HSet(ctx, key, "description", category.Description)

		redisdb.UpdateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (c *Category) Get(code string) (category *Category, err error) {

	redisClient, err := redisdb.New(redisdb.CategoryDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()
	category = &Category{}

	key := fmt.Sprintf("categories:%v", code)

	res := redisClient.HGetAll(ctx, key)

	err = res.Err()

	if err != nil {
		return
	}

	err = res.Scan(category)

	if err != nil {
		return
	}

	err = utils.InjectMap(res.Val(), category)

	if err != nil {
		return
	}

	if category.Status != status.Enabled {
		return nil, nil
	}

	return
}

func (c *Category) GetDataInfo(code string) (dataInfo *redisdb.DataInfo, err error) {

	dataInfo, err = redisdb.GetDataInfo(redisdb.CategoryDatabase, fmt.Sprintf("categories:%v", code))

	if err != nil || dataInfo == nil {
		return
	}

	return
}

func (c *Category) GetList() (categories []Category, err error) {

	redisClient, err := redisdb.New(redisdb.CategoryDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	iter := redisClient.Scan(ctx, 0, "categories:*", 0).Iterator()

	err = iter.Err()

	if err != nil {
		return
	}

	for iter.Next(ctx) {
		category, er := c.Get(strings.Replace(iter.Val(), "categories:", "", 2))

		if er != nil {
			continue
		}

		if category == nil {
			continue
		}

		categories = append(categories, *category)
	}

	return
}

func (c *Category) GetKeys() (codes []string, err error) {

	redisClient, err := redisdb.New(redisdb.CategoryDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	iter := redisClient.Scan(ctx, 0, "categories:*", 0).Iterator()

	for iter.Next(ctx) {
		codes = append(codes, strings.Replace(iter.Val(), "categories:", "", 2))
	}

	err = iter.Err()

	if err != nil {
		return
	}

	return
}

func (c *Category) Delete(code string) (err error) {

	exists, err := c.Exists(code, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Category does not exist")
	}

	redisClient, err = redisdb.New(redisdb.CategoryDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("categories:%v", code)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Disabled)

		return
	})

	return
}

func (c *Category) Restore(code string) (err error) {

	exists, err := c.Exists(code, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Category does not exist")
	}

	redisClient, err = redisdb.New(redisdb.CategoryDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("categories:%v", code)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Enabled)

		return
	})

	if err != nil {
		return
	}

	return
}
