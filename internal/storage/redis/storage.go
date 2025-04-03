package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"pc_club_server/internal/config"
	"pc_club_server/internal/lib/redis/redisEx"
	"reflect"
	"strconv"
	"time"
)

type Storage struct {
	cfg *config.RedisConfig
	cl  *redis.Client
	clx *redisEx.Client
}

func New(ctx context.Context, cfg *config.RedisConfig) (*Storage, error) {
	const op = "storage.redis.New"

	cl := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.Database,
	})

	_, err := cl.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to connect redis: %w", op, err)
	}

	clx := redisEx.New(cl)
	clx.TranslateSetMap = map[string]func(value reflect.Value) (any, bool, error){
		"time.Time": func(value reflect.Value) (any, bool, error) {
			unix := value.MethodByName("Unix").Call(nil)[0]
			res := value.MethodByName("IsZero").Call(nil)[0]
			return unix.Interface(), !res.Bool(), nil
		},
	}
	clx.TranslateGetMap = map[string]func(val string) (any, error){
		"time.Time": func(val string) (any, error) {
			unix, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, err
			}
			return time.Unix(unix, 0), nil
		},
	}

	return &Storage{
		cfg: cfg,
		cl:  cl,
		clx: clx,
	}, nil
}
