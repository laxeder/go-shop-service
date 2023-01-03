package redisdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	stt "github.com/laxeder/go-shop-service/pkg/shared/status"
)

var log = logger.New()

var Client *redis.Client
var Port string = "6379"
var ConnectString string = "redis"

/**
* @param: Addr {string}: o nome redis é a referencia para o container do redis. Deve-se trocar 127.0.0.7 pelo nome do container
 */
func New(dababaseName Nodedatabase) (Client *redis.Client, err error) {
	ctx := context.Background()
	Client = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", ConnectString, Port), Password: "", DB: int(dababaseName)})
	_, err = Client.Ping(ctx).Result()
	return
}

func Health() (string, error) {
	ctx := context.Background()
	client, err := New(UserDatabase)
	if err != nil {
		now := time.Now().UTC()
		var down string = fmt.Sprintf("DOWN: %v", now.Format(time.RFC3339))
		return "", errors.New(down)
	}

	return fmt.Sprintf("[REDIS] UP %v", client.Time(ctx).String()), nil
}

func GetStatus(database Nodedatabase, key string) (status stt.Status, err error) {

	ctx := context.Background()
	status = stt.Disabled
	err = nil

	redisClient, err := New(database)
	if err != nil {
		return
	}

	res := redisClient.HMGet(ctx, key, "status")
	err = res.Err()

	if err != nil {
		return
	}

	databaseStatus := &stt.DatabaseSatatus{}

	err = res.Scan(databaseStatus)
	if err != nil {
		return
	}

	if databaseStatus.Status == "" {
		return
	}

	status = databaseStatus.Status

	return
}

func Exists(database Nodedatabase, key string, field string) (exists bool, err error) {

	ctx := context.Background()
	exists = false
	err = nil

	redisClient, err := New(database)
	if err != nil {
		return
	}

	res := redisClient.HExists(ctx, key, field)
	err = res.Err()

	if err != nil {
		return
	}

	exists = res.Val()

	return
}
