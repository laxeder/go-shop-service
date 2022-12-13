package address

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var redisClient *redis.Client

func Repository() *Address {
	return &Address{}
}

func (u *Address) Save(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(address.Document)

	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	key := fmt.Sprintf("address:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", address.Uuid)
		rdb.HSet(ctx, key, "number", address.Number)
		rdb.HSet(ctx, key, "zip", address.Zip)
		rdb.HSet(ctx, key, "street", address.Street)
		rdb.HSet(ctx, key, "document", document)
		rdb.HSet(ctx, key, "neighborhood", address.Neighborhood)
		rdb.HSet(ctx, key, "city", address.City)
		rdb.HSet(ctx, key, "state", address.State)
		rdb.HSet(ctx, key, "status", string(address.Status))
		rdb.HSet(ctx, key, "created_at", address.CreatedAt)
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir o endereço com documento %v no redis.", document)
		return err
	}

	return nil
}

func (u *Address) Update(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(address.Document)

	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	key := fmt.Sprintf("address:%v", document)
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
		log.Error().Err(err).Msgf("Não foi possível atualizar o endereço com documento %v no redis.", document)
		return err
	}

	return nil
}

func (u *Address) Delete(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(address.Document)

	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	address.Status = Disabled

	key := fmt.Sprintf("address:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(address.Status))
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o endereço com documento %v no redis.", document)
		return err
	}

	return nil
}

func (u *Address) Restore(address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(address.Document)

	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	address.Status = Enabled

	key := fmt.Sprintf("address:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(address.Status))
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível restaurar o endereço com documento %v no redis.", document)
		return err
	}

	return nil
}

func (u *Address) GetDocument(document string) (address *Address, err error) {

	var log = logger.New()

	address = &Address{}
	err = nil

	ctx := context.Background()
	document = str.DocumentPad(document)

	redisClient, err := redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return nil, err
	}

	key := fmt.Sprintf("address:%v", document)
	res := redisClient.HMGet(ctx, key, "uuid", "document", "status")
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o endereço com documento: %v.", document)
		return u, err
	}

	err = res.Scan(address)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um endereço válido para o documento %v.", document)
		return nil, err
	}

	if address.Status == Disabled {
		address = &Address{Status: Disabled}
		return address, nil
	}

	return address, nil
}

func (u *Address) GetByDocument(document string) (address *Address, err error) {

	var log = logger.New()

	address = &Address{}
	err = nil

	ctx := context.Background()
	document = str.DocumentPad(document)

	redisClient, err := redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return nil, err
	}

	key := fmt.Sprintf("address:%v", document)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o endereço com documento: %v.", document)
		return nil, err
	}

	err = res.Scan(address)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um endereço válido para o documento %v.", document)
		return nil, err
	}

	if address.Status == Disabled {
		address = &Address{Status: Disabled}
		return address, nil
	}

	return address, nil
}

func (u *Address) SaveDocument(dcm string, address *Address) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(address.Document)

	redisClient, err = redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return err
	}

	key := fmt.Sprintf("address:%v", dcm)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "number", address.Number)
		rdb.HSet(ctx, key, "zip", address.Zip)
		rdb.HSet(ctx, key, "street", address.Street)
		rdb.HSet(ctx, key, "neighborhood", address.Neighborhood)
		rdb.HSet(ctx, key, "city", address.City)
		rdb.HSet(ctx, key, "updated_at", address.UpdatedAt)
		return err
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o endereço com documento %v no redis.", document)
		return err
	}

	return nil
}

func (u *Address) GetList() (addresss []Address, err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.AddressDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AddressDatabase)
		return nil, err
	}

	iter := redisClient.Scan(ctx, 0, "address:*", 0).Iterator()
	for iter.Next(ctx) {
		document := iter.Val()
		address, uErr := u.GetByDocument(document)
		if uErr != nil {
			continue
		}

		if address.Status == Disabled {
			continue
		}

		addresss = append(addresss, *address)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar os endereços do banco de dados.")
		return nil, err
	}

	return addresss, nil
}
