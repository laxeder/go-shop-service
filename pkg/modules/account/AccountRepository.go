package account

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

func Repository() *Account {
	return &Account{}
}

func (a *Account) Exists(uuid string, ignoreStatus bool) (bool, error) {

	accountData, err := redisdb.GetDataInfo(redisdb.AccountDatabase, fmt.Sprintf("accounts:%v", uuid))

	if err != nil {
		return false, err
	}

	if accountData == nil {
		return false, nil
	}

	if !ignoreStatus && accountData.Status != status.Enabled {
		return false, nil
	}

	return true, nil
}

func (a *Account) Save(account *Account) (err error) {

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("accounts:%v", account.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "uuid", account.Uuid)
		rdb.HSet(ctx, key, "nickname", account.Nickname)
		rdb.HSet(ctx, key, "profession", account.Profession)
		rdb.HSet(ctx, key, "rg", account.RG)
		rdb.HSet(ctx, key, "birthday", account.Birthday)
		rdb.HSet(ctx, key, "gender", string(account.Gender))

		redisdb.CreateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (a *Account) Update(account *Account) (err error) {

	exists, err := a.Exists(account.Uuid, false)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Account does not exist")
	}

	redisClient, err = redisdb.New(redisdb.AccountDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("accounts:%v", account.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "nickname", account.Nickname)
		rdb.HSet(ctx, key, "profession", account.Profession)
		rdb.HSet(ctx, key, "rg", account.RG)
		rdb.HSet(ctx, key, "birthday", account.Birthday)
		rdb.HSet(ctx, key, "gender", string(account.Gender))

		redisdb.UpdateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (a *Account) Get(uid string) (account *Account, err error) {

	redisClient, err := redisdb.New(redisdb.AccountDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()
	account = &Account{}

	key := fmt.Sprintf("accounts:%v", uid)

	res := redisClient.HGetAll(ctx, key)

	err = res.Err()

	if err != nil {
		return
	}

	err = res.Scan(account)

	if err != nil {
		return
	}

	err = utils.InjectMap(res.Val(), account)

	if err != nil {
		return
	}

	if account.Status != status.Enabled {
		return nil, nil
	}

	return
}

func (a *Account) GetList() (accounts []Account, err error) {

	redisClient, err := redisdb.New(redisdb.AccountDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	iter := redisClient.Scan(ctx, 0, "accounts:*", 0).Iterator()

	err = iter.Err()

	if err != nil {
		return
	}

	for iter.Next(ctx) {
		account, er := a.Get(strings.Replace(iter.Val(), "accounts:", "", 2))

		if er != nil {
			continue
		}

		if account == nil {
			continue
		}

		accounts = append(accounts, *account)
	}

	return
}

func (a *Account) GetDataInfo(uuid string) (dataInfo *redisdb.DataInfo, err error) {

	dataInfo, err = redisdb.GetDataInfo(redisdb.AccountDatabase, fmt.Sprintf("accounts:%v", uuid))

	if err != nil || dataInfo == nil {
		return
	}

	return
}

func (a *Account) Delete(uuid string) (err error) {

	exists, err := a.Exists(uuid, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Account does not exist")
	}

	redisClient, err = redisdb.New(redisdb.AccountDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("accounts:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Disabled)

		return
	})

	return
}

func (a *Account) Restore(uuid string) (err error) {

	exists, err := a.Exists(uuid, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("Account does not exist")
	}

	redisClient, err = redisdb.New(redisdb.AccountDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("accounts:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Enabled)

		return
	})

	if err != nil {
		return
	}

	return
}
