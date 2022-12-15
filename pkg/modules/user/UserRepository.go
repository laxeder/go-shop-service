package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var redisClient *redis.Client

func Repository() *User {
	return &User{}
}

func MarshalBinary(str []string) (data []byte) {
	var log = logger.New()

	data, err := json.Marshal(str)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tranformar array em bytes %s", str)
	}

	return
}

func (u *User) Save(user *User) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(user.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return err
	}

	accounts := MarshalBinary(user.Accounts)
	adresses := MarshalBinary(user.Adresses)

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", user.Uuid)
		rdb.HSet(ctx, key, "full_name", user.Fullname)
		rdb.HSet(ctx, key, "first_name", user.FirstName)
		rdb.HSet(ctx, key, "last_name", user.LastName)
		rdb.HSet(ctx, key, "document", document)
		rdb.HSet(ctx, key, "email", user.Email)
		rdb.HSet(ctx, key, "telephone", user.Telephone)
		rdb.HSet(ctx, key, "password", user.Password)
		rdb.HSet(ctx, key, "salt", user.Salt)
		rdb.HSet(ctx, key, "status", string(user.Status))
		rdb.HSet(ctx, key, "adresses", adresses)
		rdb.HSet(ctx, key, "accounts", accounts)
		rdb.HSet(ctx, key, "created_at", user.CreatedAt)
		rdb.HSet(ctx, key, "updated_at", user.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir o usuário com uuid %v no redis.", document)
		return err
	}

	return nil
}

func (u *User) Update(user *User) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(user.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return err
	}

	accounts := MarshalBinary(user.Accounts)
	adresses := MarshalBinary(user.Adresses)

	key := fmt.Sprintf("users:%v", user.Uuid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "full_name", user.Fullname)
		rdb.HSet(ctx, key, "first_name", user.FirstName)
		rdb.HSet(ctx, key, "last_name", user.LastName)
		rdb.HSet(ctx, key, "email", user.Email)
		rdb.HSet(ctx, key, "telephone", user.Telephone)
		rdb.HSet(ctx, key, "adresses", adresses)
		rdb.HSet(ctx, key, "accounts", accounts)
		rdb.HSet(ctx, key, "updated_at", user.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o usuário com uuid %v no redis.", document)
		return err
	}

	return nil
}

func (u *User) Delete(user *User) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(user.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return err
	}

	user.Status = Disabled

	key := fmt.Sprintf("users:%v", user.Uuid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(user.Status))
		rdb.HSet(ctx, key, "updated_at", user.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o usuário com uuid %v no redis.", document)
		return err
	}

	return nil
}

func (u *User) Restore(user *User) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(user.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return err
	}

	user.Status = Enabled

	key := fmt.Sprintf("users:%v", user.Uuid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(user.Status))
		rdb.HSet(ctx, key, "updated_at", user.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível retaurar o usuário com uuid %v no redis.", document)
		return err
	}

	return nil
}

func (u *User) SavePassowrd(user *User) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(user.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return err
	}

	key := fmt.Sprintf("users:%v", user.Uuid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "password", user.Password)
		rdb.HSet(ctx, key, "salt", user.Salt)
		rdb.HSet(ctx, key, "updated_at", user.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o usuário com uuid %v no redis.", document)
		return err
	}

	return nil
}

func (u *User) SaveDocument(user *User) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(user.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return err
	}

	key := fmt.Sprintf("users:%v", user.Uuid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "document", document)
		rdb.HSet(ctx, key, "updated_at", user.UpdatedAt)
		return err
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o usuário com uuid %v no redis.", document)
		return err
	}

	return nil
}

func (u *User) GetPasswordByUuid(uuid string) (user *User, err error) {

	var log = logger.New()

	user = &User{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return nil, err
	}

	key := fmt.Sprintf("users:%v", uuid)
	res := redisClient.HMGet(ctx, key, "uid", "uuid", "password", "salt")
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o usuário com uuid: %v.", uuid)
		return nil, err
	}

	err = res.Scan(user)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um usuário válido com o uuid %v.", uuid)
		return nil, err
	}

	if user.Status == Disabled {
		user = &User{Status: Disabled}
		return user, nil
	}

	return user, nil
}

func (u *User) GetUuid(uuid string) (user *User, err error) {

	var log = logger.New()

	user = &User{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return nil, err
	}

	key := fmt.Sprintf("users:%v", uuid)
	res := redisClient.HMGet(ctx, key, "uid", "uuid", "status")

	err = res.Err()

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o usuário com uuid: %v.", uuid)
		return u, err
	}

	err = res.Scan(user)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um usuário válido para o uuid %v.", uuid)
		return nil, err
	}

	if user.Status == Disabled {
		user = &User{Status: Disabled}
		return user, nil
	}

	return user, nil
}

func (u *User) GetByUuid(uuid string) (user *User, err error) {

	var log = logger.New()

	user = &User{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return nil, err
	}

	key := fmt.Sprintf("users:%v", uuid)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o usuário com uuid: %v.", uuid)
		return nil, err
	}

	err = res.Scan(user)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um usuário válido para o uuid %v.", uuid)
		return nil, err
	}

	if user.Status == Disabled {
		user = &User{Status: Disabled}
		return user, nil
	}

	// ? esses campos não podem ficar expostos
	user.Password = ""
	user.ConfirmPassword = ""
	user.Salt = ""

	return user, nil
}

func (u *User) GetList() (users []User, err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return nil, err
	}

	iter := redisClient.Scan(ctx, 0, "users:*", 0).Iterator()
	for iter.Next(ctx) {
		uuid := strings.Replace(iter.Val(), "users:", "", 2)
		user, uErr := u.GetByUuid(uuid)

		if uErr != nil {
			continue
		}

		if user.Status == Disabled {
			continue
		}

		user.Password = ""
		user.ConfirmPassword = ""
		user.Salt = ""

		users = append(users, *user)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar os usuários do banco de dados. %v", err)
		return nil, err
	}

	return users, nil
}
