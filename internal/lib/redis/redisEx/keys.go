package redisEx

import (
	"context"
)

func (c *Client) Keys(ctx context.Context, pattern string) ([]string, error) {
	keys, _, err := c.cl.Scan(ctx, 0, pattern, 0).Result()
	if err != nil {
		return nil, &Error{
			Code:    ErrRedisErrorCode,
			Message: "failed to scan redis keys to delete",
			Err:     err,
		}
	}
	if len(keys) == 0 {
		return nil, ErrNil
	}
	return keys, nil
}
