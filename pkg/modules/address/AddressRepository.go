package address

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

func Repository() *Address {
	return &Address{}
}

func (a *Address) Exists(uuid string, uid string, ignoreStatus bool) (bool, error) {

	addressData, err := redisdb.GetDataInfo(redisdb.AddressDatabase, fmt.Sprintf("adresses:%v:%v", uuid, uid))

	if err != nil {
		return false, err
	}

	if addressData == nil {
		return false, nil
	}

	if !ignoreStatus && addressData.Status != status.Enabled {
		return false, nil
	}

	return true, nil
}

func (a *Address) Save(address *Address) (err error) {

	redisClient, err = redisdb.New(redisdb.AddressDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("adresses:%v:%v", address.Uuid, address.Uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "uuid", address.Uuid)
		rdb.HSet(ctx, key, "uid", address.Uid)
		rdb.HSet(ctx, key, "number", address.Number)
		rdb.HSet(ctx, key, "zip", address.Zip)
		rdb.HSet(ctx, key, "street", address.Street)
		rdb.HSet(ctx, key, "neighborhood", address.Neighborhood)
		rdb.HSet(ctx, key, "city", address.City)
		rdb.HSet(ctx, key, "state", address.State)

		redisdb.CreateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (a *Address) Update(address *Address) (err error) {

	exists, err := a.Exists(address.Uuid, address.Uid, false)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Address does not exist")
	}

	redisClient, err = redisdb.New(redisdb.AddressDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("adresses:%v:%v", address.Uuid, address.Uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "number", address.Number)
		rdb.HSet(ctx, key, "zip", address.Zip)
		rdb.HSet(ctx, key, "street", address.Street)
		rdb.HSet(ctx, key, "neighborhood", address.Neighborhood)
		rdb.HSet(ctx, key, "city", address.City)
		rdb.HSet(ctx, key, "state", address.State)

		redisdb.UpdateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (a *Address) Get(uuid string, uid string) (address *Address, err error) {

	redisClient, err := redisdb.New(redisdb.AddressDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()
	address = &Address{}

	key := fmt.Sprintf("adresses:%v:%v", address.Uuid, address.Uid)

	res := redisClient.HGetAll(ctx, key)

	err = res.Err()

	if err != nil {
		return
	}

	err = res.Scan(address)

	if err != nil {
		return
	}

	err = utils.InjectMap(res.Val(), address)

	if err != nil {
		return
	}

	if address.Status != status.Enabled {
		return nil, nil
	}

	return
}

func (a *Address) GetDataInfo(uuid string, uid string) (dataInfo *redisdb.DataInfo, err error) {

	dataInfo, err = redisdb.GetDataInfo(redisdb.AddressDatabase, fmt.Sprintf("adresses:%v:%v", uuid, uid))

	if err != nil || dataInfo == nil {
		return
	}

	return
}

func (a *Address) GetList(uuid string) (adresses []Address, err error) {

	redisClient, err := redisdb.New(redisdb.AddressDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("adresses:%v:*", uuid)

	iter := redisClient.Scan(ctx, 0, key, 0).Iterator()

	err = iter.Err()

	if err != nil {
		return
	}

	for iter.Next(ctx) {
		address, er := a.Get(uuid, strings.Replace(iter.Val(), key, "", 2))

		if er != nil {
			continue
		}

		if address == nil {
			continue
		}

		adresses = append(adresses, *address)
	}

	return
}

func (a *Address) Delete(uuid string, uid string) (err error) {

	exists, err := a.Exists(uuid, uid, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Address does not exist")
	}

	redisClient, err = redisdb.New(redisdb.AddressDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("adresses:%v:%v", uuid, uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Disabled)

		return
	})

	return
}

func (a *Address) Restore(uuid string, uid string) (err error) {

	exists, err := a.Exists(uuid, uid, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Address does not exist")
	}

	redisClient, err = redisdb.New(redisdb.AddressDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("adresses:%v:%v", uuid, uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Enabled)

		return
	})

	if err != nil {
		return
	}

	return
}
