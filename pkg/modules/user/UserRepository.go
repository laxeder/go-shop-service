package user

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

func Repository() *User {
	return &User{}
}

func (u *User) Exists(uuid string, ignoreStatus bool) (bool, error) {

	userData, err := redisdb.GetDataInfo(redisdb.UserDatabase, fmt.Sprintf("users:%v", uuid))

	if err != nil {
		return false, err
	}

	if userData == nil {
		return false, nil
	}

	if !ignoreStatus && userData.Status != status.Enabled {
		return false, nil
	}

	return true, nil
}

func (u *User) Save(user *User) (err error) {

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	user.GenerateUuid()
	user.GenerateFullName()
	user.GenerateDocument()
	user.GeneratePassword()

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "uuid", user.Uuid)
		rdb.HSet(ctx, key, "full_name", user.Fullname)
		rdb.HSet(ctx, key, "first_name", user.FirstName)
		rdb.HSet(ctx, key, "last_name", user.LastName)
		rdb.HSet(ctx, key, "document", user.Document)
		rdb.HSet(ctx, key, "email", user.Email)
		rdb.HSet(ctx, key, "telephone", user.Telephone)
		rdb.HSet(ctx, key, "password", user.Password)
		rdb.HSet(ctx, key, "salt", user.Salt)

		redisdb.CreateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) SavePassword(user *User) (err error) {

	exists, err := u.Exists(user.Uuid, false)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("User does not exist")
	}

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	user.GeneratePassword()

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "password", user.Password)
		rdb.HSet(ctx, key, "salt", user.Salt)

		redisdb.UpdateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) SaveDocument(user *User) (err error) {

	exists, err := u.Exists(user.Uuid, false)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("User does not exist")
	}

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	user.GenerateDocument()

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "document", user.Document)

		redisdb.UpdateDataInfo(rdb, ctx, key)

		return err
	})

	if err != nil {
		return
	}

	return
}

func (u *User) Update(user *User) (err error) {

	exists, err := u.Exists(user.Uuid, false)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("User does not exist")
	}

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	user.GenerateFullName()

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		rdb.HSet(ctx, key, "full_name", user.Fullname)
		rdb.HSet(ctx, key, "first_name", user.FirstName)
		rdb.HSet(ctx, key, "last_name", user.LastName)
		rdb.HSet(ctx, key, "email", user.Email)
		rdb.HSet(ctx, key, "telephone", user.Telephone)

		redisdb.UpdateDataInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) Get(uuid string) (user *User, err error) {

	redisClient, err := redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()
	user = &User{}

	key := fmt.Sprintf("users:%v", uuid)

	res := redisClient.HGetAll(ctx, key)

	err = res.Err()

	if err != nil {
		return
	}

	err = res.Scan(user)

	if err != nil {
		return
	}

	err = utils.InjectMap(res.Val(), user)

	if err != nil {
		return
	}

	if user.Status != status.Enabled {
		return nil, nil
	}

	//? esses campos não podem ficar expostos
	user.Password = ""
	user.Salt = ""

	return
}

func (u *User) GetDataInfo(uuid string) (dataInfo *redisdb.DataInfo, err error) {

	dataInfo, err = redisdb.GetDataInfo(redisdb.UserDatabase, fmt.Sprintf("users:%v", uuid))

	if err != nil || dataInfo == nil {
		return
	}

	return
}

func (u *User) GetPassword(uuid string) (user *User, err error) {

	redisClient, err := redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()
	user = &User{}

	key := fmt.Sprintf("users:%v", uuid)

	res := redisClient.HGetAll(ctx, key)

	err = res.Err()

	if err != nil {
		return
	}

	err = res.Scan(user)

	if err != nil {
		return
	}

	err = utils.InjectMap(res.Val(), user)

	if err != nil {
		return
	}

	if user.Status != status.Enabled {
		return nil, nil
	}

	return
}

func (u *User) GetByEmail(email string) (user *User, err error) {

	users, err := u.GetList()

	if err != nil {
		return
	}

	for _, usr := range users {
		if usr.Email == email {
			user = &usr
			break
		}
	}

	return
}

func (u *User) GetList() (users []User, err error) {

	redisClient, err := redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	iter := redisClient.Scan(ctx, 0, "users:*", 0).Iterator()

	err = iter.Err()

	if err != nil {
		return
	}

	for iter.Next(ctx) {
		user, er := u.Get(strings.Replace(iter.Val(), "users:", "", 2))

		if er != nil {
			continue
		}

		if user == nil {
			continue
		}

		users = append(users, *user)
	}

	return
}

func (u *User) Delete(uuid string) (err error) {

	exists, err := u.Exists(uuid, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("User does not exist")
	}

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("users:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Disabled)

		return
	})

	return
}

func (u *User) Restore(uuid string) (err error) {

	exists, err := u.Exists(uuid, true)

	if err != nil {
		return
	}

	if !exists {
		return fmt.Errorf("User does not exist")
	}

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	ctx := context.Background()

	key := fmt.Sprintf("users:%v", uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {

		redisdb.UpdateDataStatus(rdb, ctx, key, status.Enabled)

		return
	})

	if err != nil {
		return
	}

	return
}
