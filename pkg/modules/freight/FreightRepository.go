package freight

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
)

var redisClient *redis.Client

func Repository() *Freight {
	return &Freight{}
}

func (f *Freight) Save(freight *Freight) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uid := freight.Uid

	redisClient, err = redisdb.New(redisdb.FreightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.FreightDatabase)
		return err
	}

	key := fmt.Sprintf("freights:%v", uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uid", uid)
		rdb.HSet(ctx, key, "zipcode_sender", freight.ZipcodeSender)
		rdb.HSet(ctx, key, "zipcode_receiver", freight.ZipcodeReceiver)
		rdb.HSet(ctx, key, "type", string(freight.Type))
		rdb.HSet(ctx, key, "price", freight.Price)
		rdb.HSet(ctx, key, "weight", freight.Weight)
		rdb.HSet(ctx, key, "heigth", freight.Heigth)
		rdb.HSet(ctx, key, "width", freight.Width)
		rdb.HSet(ctx, key, "lenght", freight.Lenght)
		rdb.HSet(ctx, key, "status", string(freight.Status))
		rdb.HSet(ctx, key, "updated_at", freight.UpdatedAt)
		rdb.HSet(ctx, key, "created_at", freight.CreatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir o frete com uid %v no redis.", uid)
		return err
	}

	return nil
}

func (f *Freight) Update(freight *Freight) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uid := freight.Uid

	redisClient, err = redisdb.New(redisdb.FreightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.FreightDatabase)
		return err
	}

	key := fmt.Sprintf("freights:%v", uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "zipcode_sender", freight.ZipcodeSender)
		rdb.HSet(ctx, key, "zipcode_receiver", freight.ZipcodeReceiver)
		rdb.HSet(ctx, key, "type", string(freight.Type))
		rdb.HSet(ctx, key, "price", freight.Price)
		rdb.HSet(ctx, key, "weight", freight.Weight)
		rdb.HSet(ctx, key, "heigth", freight.Heigth)
		rdb.HSet(ctx, key, "width", freight.Width)
		rdb.HSet(ctx, key, "lenght", freight.Lenght)
		rdb.HSet(ctx, key, "status", string(freight.Status))
		rdb.HSet(ctx, key, "updated_at", freight.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o frete com uid %v no redis.", uid)
		return err
	}

	return nil
}

func (f *Freight) Delete(freight *Freight) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uid := freight.Uid

	redisClient, err = redisdb.New(redisdb.FreightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.FreightDatabase)
		return err
	}

	freight.Status = Disabled

	key := fmt.Sprintf("freights:%v", uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(freight.Status))
		rdb.HSet(ctx, key, "updated_at", freight.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o frete com uid %v no redis.", uid)
		return err
	}

	return nil
}

func (f *Freight) Restore(freight *Freight) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	uid := freight.Uid

	redisClient, err = redisdb.New(redisdb.FreightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.FreightDatabase)
		return err
	}

	freight.Status = Enabled

	key := fmt.Sprintf("freights:%v", uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(freight.Status))
		rdb.HSet(ctx, key, "updated_at", freight.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível retaurar o frete com uid %v no redis.", uid)
		return err
	}

	return nil
}

func (f *Freight) GetUid(uid string) (freight *Freight, err error) {

	var log = logger.New()

	freight = &Freight{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.FreightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.FreightDatabase)
		return nil, err
	}

	key := fmt.Sprintf("freights:%v", uid)
	res := redisClient.HMGet(ctx, key, "uid", "uid", "status")

	err = res.Err()

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o frete com uid: %v.", uid)
		return f, err
	}

	err = res.Scan(freight)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um frete válido para o uid %v.", uid)
		return nil, err
	}

	if freight.Status == Disabled {
		freight = &Freight{Status: Disabled}
		return freight, nil
	}

	return freight, nil
}

func (f *Freight) GetByUid(uid string) (freight *Freight, err error) {

	var log = logger.New()

	freight = &Freight{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.FreightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.FreightDatabase)
		return nil, err
	}

	key := fmt.Sprintf("freights:%v", uid)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o frete com uid: %v.", uid)
		return nil, err
	}

	err = res.Scan(freight)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um frete válido para o uid %v.", uid)
		return nil, err
	}

	if freight.Status == Disabled {
		freight = &Freight{Status: Disabled}
		return freight, nil
	}

	return freight, nil
}

func (f *Freight) GetList() (freights []Freight, err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.FreightDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.FreightDatabase)
		return nil, err
	}

	iter := redisClient.Scan(ctx, 0, "freights:*", 0).Iterator()
	for iter.Next(ctx) {
		uid := strings.Replace(iter.Val(), "freights:", "", 2)
		freight, uErr := f.GetByUid(uid)

		if uErr != nil {
			continue
		}

		if freight.Status == Disabled {
			continue
		}

		freights = append(freights, *freight)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar os fretes do banco de dados.")
		return nil, err
	}

	return freights, nil
}
