package auth

import (
	"context"
	"pc_club_server/internal/config"
	"time"
)

type redisProvider interface {
	Get(
		ctx context.Context,
		value any,
		tags ...any,
	) error
}

type redisOwner interface {
	SetWithTTL(
		ctx context.Context,
		ttl time.Duration,
		value any,
		tags ...any,
	) error
}

type versionProvider interface {
	RefreshVersion(
		ctx context.Context,
		uid int64,
	) (version int64, err error)
}

type versionOwner interface {
	IncRefreshVersion(
		ctx context.Context,
		uid int64,
	) (err error)
}

type Service struct {
	cfg             *config.AuthConfig
	redisOwner      redisOwner
	redisProvider   redisProvider
	versionProvider versionProvider
	versionOwner    versionOwner
}

const (
	RefreshRedisBlackListName = "refresh_black_list_exp"
	AccessRedisBlackListName  = "access_black_list_exp"
)

func NewService(
	cfg *config.AuthConfig,
	redisOwner redisOwner,
	redisProvider redisProvider,
	VersionOwner versionOwner,
	VersionProvider versionProvider,
) *Service {
	return &Service{
		cfg:             cfg,
		redisOwner:      redisOwner,
		redisProvider:   redisProvider,
		versionOwner:    VersionOwner,
		versionProvider: VersionProvider,
	}
}
