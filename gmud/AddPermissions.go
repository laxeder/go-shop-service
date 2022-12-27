package gmud

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

var redisClient *redis.Client

func AddPermissions() {
	var log = logger.New()

	var permissions []string = []string{
		"getProduct",
		"getCatgeory",
		"getUser",
		"getAccount",
		"getAddress",
	}

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.PermissionDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.PermissionDatabase)
		return
	}

	users, _ := user.Repository().GetList()

	for _, userData := range users {
		key := fmt.Sprintf("permissions:%v", userData.Uuid)

		_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
			rdb.HSet(ctx, key, "permissions", user.MarshalBinary(permissions))

			return nil
		})

	}

	//TODO: A cada nova permissão atualizar permissões do usuário que não a contem
	//TODO: implementar script mudar permissões do banco de dados
	//TODO: adicionar novas permissões (não só sobescrever)
	//TODO: não repetir permissões

	// read := &utils.ReadFile{}
	// scriptFile := read.Run("./gmud/permissions.rd")

	// spew.Dump(scriptFile)

	// vals, err := redisClient.Eval(ctx, scriptFile, []string{"users:abc123"}, "hello").Result()
	// if err != nil {
	// 	log.Error().Err(err).Msgf("Erro ao executar script (%v)", scriptFile)
	// 	return
	// }
}
