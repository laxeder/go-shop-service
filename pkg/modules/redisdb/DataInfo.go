package redisdb

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

type DataInfo struct {
	Status    status.Status `json:"status,omitempty" redis:"status,omitempty"`
	CreatedAt string        `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt string        `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func CreateDataInfo(rdb redis.Pipeliner, ctx context.Context, key string) {
	rdb.HSet(ctx, key, "status", string(status.Enabled))
	rdb.HSet(ctx, key, "created_at", date.NowUTC())
	rdb.HSet(ctx, key, "updated_at", date.NowUTC())
}

func UpdateDataInfo(rdb redis.Pipeliner, ctx context.Context, key string) {
	rdb.HSet(ctx, key, "updated_at", date.NowUTC())
}

func UpdateDataStatus(rdb redis.Pipeliner, ctx context.Context, key string, status status.Status) {
	rdb.HSet(ctx, key, "status", string(status))
	rdb.HSet(ctx, key, "updated_at", date.NowUTC())
}

func GetDataInfo(database Nodedatabase, key string) (dataInfo *DataInfo, err error) {

	ctx := context.Background()
	dataInfo = &DataInfo{}

	redisClient, err := New(database)

	if err != nil {
		return nil, err
	}

	res := redisClient.HMGet(ctx, key, "status", "created_at", "updated_at")

	err = res.Err()

	if err != nil {
		return nil, err
	}

	if len(res.Val()) == 0 {
		return nil, nil
	}

	err = res.Scan(dataInfo)

	if err != nil {
		return nil, err
	}

	return
}
