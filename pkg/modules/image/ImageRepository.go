package image

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var redisClient *redis.Client

func Repository() *Image {
	return &Image{}
}

func (i *Image) Save(image *Image) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(image.Document)

	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", image.Uuid)
		rdb.HSet(ctx, key, "document", document)
		rdb.HSet(ctx, key, "size", image.Size)
		rdb.HSet(ctx, key, "mime_type", image.MimeType)
		rdb.HSet(ctx, key, "Name", image.Name)
		rdb.HSet(ctx, key, "base64", image.Base64)
		rdb.HSet(ctx, key, "status", string(image.Status))
		rdb.HSet(ctx, key, "created_at", image.CreatedAt)
		rdb.HSet(ctx, key, "updated_at", image.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir uma imagem para o documento %v no redis.", document)
		return
	}

	return
}

func (i *Image) Update(image *Image) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(image.Document)

	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "size", image.Size)
		rdb.HSet(ctx, key, "mime_type", image.MimeType)
		rdb.HSet(ctx, key, "Name", image.Name)
		rdb.HSet(ctx, key, "base64", image.Base64)
		rdb.HSet(ctx, key, "status", string(image.Status))
		rdb.HSet(ctx, key, "updated_at", image.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar a imagem para o documento %v no redis.", document)
		return
	}

	return
}

func (i *Image) Delete(image *Image) (err error) {

	var log = logger.New()

	ctx := context.Background()
	document := str.DocumentPad(image.Document)

	err = nil

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	image.Status = Disabled

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(image.Status))
		rdb.HSet(ctx, key, "updated_at", image.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o usuário com documento %v no redis.", document)
		return
	}

	return
}

func (i *Image) Restore(image *Image) (err error) {

	var log = logger.New()

	ctx := context.Background()
	err = nil

	document := str.DocumentPad(image.Document)

	redisClient, err = redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	image.Status = Enabled

	key := fmt.Sprintf("images:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(image.Status))
		rdb.HSet(ctx, key, "updated_at", image.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível retaurar a imagem para o documento %v no redis.", document)
		return
	}

	return
}

func (i *Image) GetByDocument(document string) (image *Image, err error) {

	var log = logger.New()

	image = &Image{}

	ctx := context.Background()
	document = str.DocumentPad(document)

	redisClient, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	key := fmt.Sprintf("images:%v", document)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar a imagem para o documento: %v.", document)
		return
	}

	err = res.Scan(image)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear uma imagem válida para o documento %v.", document)
		return
	}

	if image.Status == Disabled {
		image = &Image{Status: Disabled}
		return
	}

	return
}

func (i *Image) GetList() (images []Image, err error) {

	var log = logger.New()

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.UserDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.UserDatabase)
		return
	}

	iter := redisClient.Scan(ctx, 0, "images:*", 0).Iterator()
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
		log.Error().Err(err).Msgf("Não foi possível listar as imagens do banco de dados. %v", err)
		return
	}

	return
}
