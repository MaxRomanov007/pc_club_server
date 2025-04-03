package redisEx

import (
	"github.com/redis/go-redis/v9"
	"reflect"
)

type Client struct {
	cl              *redis.Client
	TranslateSetMap map[string]func(value reflect.Value) (any, bool, error)
	TranslateGetMap map[string]func(val string) (any, error)
}

func New(cl *redis.Client) *Client {
	return &Client{
		cl: cl,
	}
}
