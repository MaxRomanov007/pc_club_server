package redisEx

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"pc_club_server/internal/lib/redis/redisKey"
)

func (c *Client) Del(ctx context.Context, val any, tags ...any) error {
	keys, err := c.Keys(ctx, redisKey.Key(val, tags...)+".*")
	if err != nil {
		return err
	}
	_, err = c.cl.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			c.cl.Del(ctx, key)
		}
		return nil
	})
	if err != nil {
		return &Error{
			Code:    ErrRedisErrorCode,
			Message: "failed to execute delete pipeline",
			Err:     err,
		}
	}
	return nil
}
