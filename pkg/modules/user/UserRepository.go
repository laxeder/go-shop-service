package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

var redisClient *redis.Client

func Repository() *User {
	return &User{}
}

func (u *User) Save(user *User) (err error) {

	ctx := context.Background()
	err = nil

	user.GenerateUuid()
	user.GenerateFullName()
	user.GenerateDocument()
	user.GeneratePassword()

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

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

		redisdb.CreateItemInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) Update(user *User) (err error) {

	ctx := context.Background()
	err = nil

	user.GenerateFullName()

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {
		rdb.HSet(ctx, key, "full_name", user.Fullname)
		rdb.HSet(ctx, key, "first_name", user.FirstName)
		rdb.HSet(ctx, key, "last_name", user.LastName)
		rdb.HSet(ctx, key, "email", user.Email)
		rdb.HSet(ctx, key, "telephone", user.Telephone)

		redisdb.UpdateItemInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) Delete(user *User) (err error) {

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {
		redisdb.UpdateItemStatus(rdb, ctx, key, status.Disabled)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) Restore(user *User) (err error) {

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {
		redisdb.UpdateItemStatus(rdb, ctx, key, status.Enabled)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) SavePassowrd(user *User) (err error) {

	ctx := context.Background()
	err = nil

	user.GeneratePassword()

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {
		rdb.HSet(ctx, key, "password", user.Password)
		rdb.HSet(ctx, key, "salt", user.Salt)

		redisdb.UpdateItemInfo(rdb, ctx, key)

		return
	})

	if err != nil {
		return
	}

	return
}

func (u *User) SaveDocument(user *User) (err error) {

	ctx := context.Background()
	err = nil

	user.GenerateDocument()

	redisClient, err = redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) (err error) {
		rdb.HSet(ctx, key, "document", user.Document)

		redisdb.UpdateItemInfo(rdb, ctx, key)

		return err
	})

	if err != nil {
		return
	}

	return
}

func (u *User) GetPassword(uuid string) (user *User, err error) {

	ctx := context.Background()
	user = &User{}
	err = nil

	redisClient, err := redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

	key := fmt.Sprintf("users:%v", uuid)

	res := redisClient.HMGet(ctx, key, "uuid", "password", "salt")

	err = res.Err()

	if err != nil {
		return
	}

	err = res.Scan(user)

	if err != nil {
		return
	}

	return
}

func (u *User) GetByEmail(email string) (user *User, err error) {

	user = nil
	err = nil

	users, er := u.GetList()

	if er != nil {
		return user, er
	}

	for _, usr := range users {
		if usr.Email == email {
			user = &usr
			break
		}
	}

	return
}

func (u *User) Get(uuid string) (user *User, err error) {

	user = nil
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

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

	userInfo := &redisdb.ItemInfo{}

	err = res.Scan(userInfo)

	if err != nil {
		return
	}

	if userInfo.Status == status.Disabled {
		return nil, nil
	}

	//? esses campos não podem ficar expostos
	user.Password = ""
	user.ConfirmPassword = ""
	user.Salt = ""

	return
}

func (u *User) GetList() (users []User, err error) {

	ctx := context.Background()
	err = nil

	redisClient, err := redisdb.New(redisdb.UserDatabase)

	if err != nil {
		return
	}

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
