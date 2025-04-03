package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"pc_club_server/internal/lib/redis/redisEx"
	redis2 "pc_club_server/internal/lib/redis/redisKey"
	"time"
)

func (s *Storage) Set(ctx context.Context, value any, tags ...any) error {
	return s.SetWithTTL(ctx, s.cfg.DefaultTTL, value, tags...)
}

func (s *Storage) SetWithTTL(ctx context.Context, ttl time.Duration, value any, tags ...any) error {
	const op = "storage.redis.SetWithTTL"

	if err := s.clx.SetWithTTL(ctx, ttl, value, tags...); err != nil {
		return fmt.Errorf("%s: failed to set key in redis: %w", op, err)
	}

	return nil
}

func (s *Storage) Get(ctx context.Context, value any, tags ...any) error {
	const op = "storage.redis.Value"

	if err := s.clx.Get(ctx, value, tags...); err != nil {
		if errors.Is(err, redisEx.ErrNil) {
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return fmt.Errorf("%s: failed to get key from redis: %w", op, err)
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, value any, tags ...any) error {
	const op = "storage.redis.Delete"

	key := redis2.Key(value, tags...)
	if err := s.cl.Del(ctx, key).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return fmt.Errorf("%s: failed to del value: %w", op, err)
	}

	return nil
}
