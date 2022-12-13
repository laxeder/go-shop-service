package icp

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var redisClient *redis.Client

func Repository() *ICP {
	return &ICP{}
}

func (i *ICP) Save(icp *ICP) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(icp.Document)

	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("icps:%v", document)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", icp.Uuid)
		rdb.HSet(ctx, key, "document", icp.Document)
		rdb.HSet(ctx, key, "key_public", icp.KeyPublic)
		rdb.HSet(ctx, key, "name", icp.Name)
		rdb.HSet(ctx, key, "validate", icp.Validate)
		rdb.HSet(ctx, key, "ac_name", icp.ACName)
		rdb.HSet(ctx, key, "serial_number", icp.SerialNumber)
		rdb.HSet(ctx, key, "digital_signature", icp.DigitalSignature)
		rdb.HSet(ctx, key, "status", string(icp.Status))
		rdb.HSet(ctx, key, "created_at", icp.CreatedAt)
		rdb.HSet(ctx, key, "updated_at", icp.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir um certificado com documento %v no redis.", document)
		return
	}

	return
}

func (i *ICP) Update(icp *ICP) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(icp.Document)

	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("icps:%v", document)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "key_public", icp.KeyPublic)
		rdb.HSet(ctx, key, "name", icp.Name)
		rdb.HSet(ctx, key, "validate", icp.Validate)
		rdb.HSet(ctx, key, "ac_name", icp.ACName)
		rdb.HSet(ctx, key, "serial_number", icp.SerialNumber)
		rdb.HSet(ctx, key, "digital_signature", icp.DigitalSignature)
		rdb.HSet(ctx, key, "status", string(icp.Status))
		rdb.HSet(ctx, key, "updated_at", icp.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o certificado com documento %v no redis.", document)
		return
	}

	return
}

func (i *ICP) Delete(icp *ICP) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(icp.Document)

	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("icps:%v", document)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(icp.Status))
		rdb.HSet(ctx, key, "updated_at", icp.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o certificado com documento %v no redis.", document)
		return
	}

	return
}

func (i *ICP) Restore(icp *ICP) (err error) {

	var log = logger.New()

	ctx := context.Background()
	err = nil

	document := str.DocumentPad(icp.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("icps:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(icp.Status))
		rdb.HSet(ctx, key, "updated_at", icp.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível retaurar o certificado com documento %v no redis.", document)
		return
	}

	return
}

func (i *ICP) GetByDocument(document string) (icp *ICP, err error) {

	var log = logger.New()

	icp = &ICP{}

	ctx := context.Background()
	document = str.DocumentPad(document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("icps:%v", document)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar um certificado com o documento: %v.", document)
		return
	}

	err = res.Scan(icp)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um certificado válido com o documento %v.", document)
		return
	}

	if icp.Status == Disabled {
		icp = &ICP{Status: Disabled}
		return
	}

	return
}

func (i *ICP) GetList() (icps []ICP, err error) {

	var log = logger.New()

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	iter := redisClient.Scan(ctx, 0, "icps:*", 0).Iterator()
	for iter.Next(ctx) {
		document := iter.Val()
		icp, iErr := i.GetByDocument(document)
		if iErr != nil {
			continue
		}

		if icp.Status == Disabled {
			continue
		}

		icps = append(icps, *icp)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar os certificados do banco de dados.")
		return
	}

	return
}
