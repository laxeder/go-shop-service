package image

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var redisClientImages *redis.Client

func RepositoryImages() *Images {
	return &Images{}
}

func (i *Images) Save(images *Images) (err error) {
	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(images.Document)

	err = nil

	redisClientImages, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	images.ImagesMerge()

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClientImages.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "document", document)
		rdb.HSet(ctx, key, "images_json", images.ImagesJson)
		rdb.HSet(ctx, key, "status", string(images.Status))
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir as imagens para o documento %v no redis.", document)
		return
	}

	return
}

func (i *Images) Update(images *Images) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(i.Document)

	err = nil

	redisClientImages, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClientImages.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "document", images.Document)
		rdb.HSet(ctx, key, "images_json", images.ImagesJson)
		rdb.HSet(ctx, key, "status", string(images.Status))

		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o usuário com documento %v no redis.", document)
		return
	}

	return
}

func (i *Images) Delete(images *Images) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(images.Document)

	err = nil

	redisClientImages, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	images.Status = Disabled

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClientImages.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(images.Status))
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar as imagens para o documento %v no redis.", document)
		return
	}

	return
}

func (i *Images) Restore(images *Images) (err error) {

	var log = logger.New()

	ctx := context.Background()
	err = nil

	document := str.DocumentPad(images.Document)

	redisClientImages, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	images.Status = Enabled

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClientImages.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(images.Status))
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível retaurar o usuário com documento %v no redis.", document)
		return
	}

	return
}

func (i *Images) GetByDocument(document string) (images *Images, err error) {

	var log = logger.New()

	images = &Images{}

	ctx := context.Background()
	document = str.DocumentPad(document)

	redisClientImages, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("images:%v", document)
	res := redisClientImages.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar as imagens para o documento: %v.", document)
		return
	}

	err = res.Scan(images)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear imagens válidas para o documento %v.", document)
		return
	}

	if images.Status == Disabled {
		images = &Images{Status: Disabled}
		return
	}

	return
}

func (i *Images) GetList() (images []Images, err error) {

	var log = logger.New()

	ctx := context.Background()

	redisClientImages, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	iter := redisClientImages.Scan(ctx, 0, "images:*", 0).Iterator()
	for iter.Next(ctx) {
		document := iter.Val()
		image, iErr := i.GetByDocument(document)
		if iErr != nil {
			continue
		}

		if image.Status == Disabled {
			continue
		}

		images = append(images, *image)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar imagens do banco de dados.")
		return
	}

	return
}
