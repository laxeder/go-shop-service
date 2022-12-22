package address

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
)

var redisClient *redis.Client

func Repository() *Address {
	return &Address{}
}

func (a *Address) Save(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	key := fmt.Sprintf("adresses:%v", address.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uid", address.Uid)
		rdb.HSet(ctx, key, "uuid", address.Uuid)
		rdb.HSet(ctx, key, "number", address.Number)
		rdb.HSet(ctx, key, "zip", address.Zip)
		rdb.HSet(ctx, key, "street", address.Street)
		rdb.HSet(ctx, key, "neighborhood", address.Neighborhood)
		rdb.HSet(ctx, key, "city", address.City)
		rdb.HSet(ctx, key, "state", address.State)
		rdb.HSet(ctx, key, "status", string(address.Status))
		rdb.HSet(ctx, key, "created_at", address.CreatedAt)
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir o endereço com uid %v no redis.", address.Uid)
		return err
	}

	return nil
}

func (a *Address) Update(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	key := fmt.Sprintf("adresses:%v", address.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "number", address.Number)
		rdb.HSet(ctx, key, "zip", address.Zip)
		rdb.HSet(ctx, key, "street", address.Street)
		rdb.HSet(ctx, key, "neighborhood", address.Neighborhood)
		rdb.HSet(ctx, key, "city", address.City)
		rdb.HSet(ctx, key, "state", address.State)
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o endereço com uid %v no redis.", address.Uid)
		return err
	}

	return nil
}

func (a *Address) Delete(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	address.Status = Disabled

	key := fmt.Sprintf("adresses:%v", address.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(address.Status))
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o endereço com o uid %v no redis.", address.Uid)
		return err
	}

	return nil
}

func (a *Address) Restore(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	address.Status = Enabled

	key := fmt.Sprintf("adresses:%v", address.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(address.Status))
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível restaurar o endereço com o uid %v no redis.", address.Uid)
		return err
	}

	return nil
}

func (a *Address) GetUid(uid string) (address *Address, err error) {

	var log = logger.New()

	address = &Address{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return
	}

	key := fmt.Sprintf("adresses:%v", address.Uid)
	res := redisClient.HMGet(ctx, key, "uuid", "uid", "status")
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o endereço com o uid: %v.", uid)
		return
	}

	err = res.Scan(address)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um endereço válido para o uid %v.", uid)
		return
	}

	if address.Status == Disabled {
		address = &Address{Status: Disabled}
		return
	}

	return
}

func (a *Address) GetByUid(uid string) (address *Address, err error) {

	var log = logger.New()

	address = &Address{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return
	}

	key := fmt.Sprintf("adresses:%v", uid)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o endereço com uid: %v.", uid)
		return
	}

	err = res.Scan(address)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um endereço válido para o uid %v.", uid)
		return nil, err
	}

	if address.Status == Disabled {
		address = &Address{Status: Disabled}
		return address, nil
	}

	return address, nil
}

func (a *Address) GetList() (addresss []Address, err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return nil, err
	}

	iter := redisClient.Scan(ctx, 0, "adresses:*", 0).Iterator()
	for iter.Next(ctx) {
		uid := strings.Replace(iter.Val(), "adresses:", "", 2)
		address, aErr := a.GetByUid(uid)

		if aErr != nil {
			continue
		}

		if address.Status == Disabled {
			continue
		}

		addresss = append(addresss, *address)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar os endereços do banco de dados. %v", err)
		return nil, err
	}

	return addresss, nil
}
