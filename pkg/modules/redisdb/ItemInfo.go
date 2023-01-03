package redisdb

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/shared/status"
)

type ItemInfo struct {
	Status    status.Status `json:"status,omitempty" redis:"status,omitempty"`
	CreatedAt string        `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt string        `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func CreateItemInfo(rdb redis.Pipeliner, ctx context.Context, key string) {
	rdb.HSet(ctx, key, "status", status.Enabled)
	rdb.HSet(ctx, key, "created_at", date.NowUTC())
	rdb.HSet(ctx, key, "updated_at", date.NowUTC())
}

func UpdateItemInfo(rdb redis.Pipeliner, ctx context.Context, key string) {
	rdb.HSet(ctx, key, "updated_at", date.NowUTC())
}

func UpdateItemStatus(rdb redis.Pipeliner, ctx context.Context, key string, status status.Status) {
	rdb.HSet(ctx, key, "status", string(status))
	rdb.HSet(ctx, key, "updated_at", date.NowUTC())
}

func GetItemInfo(database Nodedatabase, key string) (itemInfo *ItemInfo, err error) {

	ctx := context.Background()
	itemInfo = &ItemInfo{}
	err = nil

	redisClient, err := New(database)

	if err != nil {
		return
	}

	res := redisClient.HMGet(ctx, key, "status", "created_at", "updated_at")

	err = res.Err()

	if err != nil {
		return
	}

	if len(res.Val()) == 0 {
		return
	}

	err = res.Scan(itemInfo)

	if err != nil {
		return
	}

	if itemInfo.Status == "" {
		itemInfo.Status = status.Disabled
	}

	return
}
