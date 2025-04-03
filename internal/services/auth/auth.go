package auth

import (
	"context"
	"errors"
	"fmt"
	"pc_club_server/internal/lib/jwt"
	"pc_club_server/internal/storage/redis"
)

type Access struct {
	ExpTime int64
}

func (a *Access) RedisKey() string {
	return AccessRedisBlackListName
}

type Refresh struct {
	ExpTime int64
}

func (a *Refresh) RedisKey() string {
	return RefreshRedisBlackListName
}

// generateTokens generates access and refresh tokens
// (1 output is access, 2 is refresh)
func (s *Service) generateTokens(
	ctx context.Context,
	uid int64,
	refreshVersion int64,
) (access string, refresh string, err error) {
	const op = "services.pcClub.auth.generateTokens"

	err = s.versionOwner.IncRefreshVersion(ctx, uid)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to update refresh version: %w", op, HandleMssqlError(err))
	}

	access, err = jwt.NewAccessToken(
		uid,
		s.cfg.Access.Secret,
		s.cfg.Access.TTL,
	)
	if err != nil {
		return "", "", err
	}

	refresh, err = jwt.NewRefreshToken(
		uid,
		refreshVersion,
		s.cfg.Refresh.Secret,
		s.cfg.Refresh.TTL,
	)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *Service) Access(
	ctx context.Context,
	accessToken string,
) (int64, error) {
	const op = "services.pcClub.auth.Access"

	claims, err := jwt.ParseToken(accessToken, s.cfg.Access.Secret)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to parse token: %w", op, TokenError(err))
	}

	var bannedToken Access
	if err = s.redisProvider.Get(ctx, &bannedToken, claims.UID); err != nil {
		if errors.Is(err, redis.ErrNotFound) {
			return claims.UID, nil
		}
		return 0, fmt.Errorf("%s: failed to find token in redis: %w", op, err)
	}

	if bannedToken.ExpTime >= claims.ExpiresAt.Unix() {
		return 0, fmt.Errorf("%s: token in black list", op)
	}

	return claims.UID, nil
}

func (s *Service) Refresh(
	ctx context.Context,
	RefreshToken string,
) (string, string, error) {
	const op = "services.pcClub.auth.Refresh"

	claims, err := jwt.ParseToken(
		RefreshToken,
		s.cfg.Refresh.Secret,
	)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to parse token: %w", op, TokenError(err))
	}

	var bannedRefresh Refresh
	if err = s.redisProvider.Get(
		ctx,
		&bannedRefresh,
		claims.UID,
	); err != nil {
		if errors.Is(err, redis.ErrNotFound) {
			access, refresh, err := s.generateTokens(ctx, claims.UID, claims.Version+1)
			if err != nil {
				return "", "", fmt.Errorf("%s: failed to generate tokens: %w", op, TokenError(err))
			}

			return access, refresh, nil
		}

		return "", "", fmt.Errorf("%s: failed to find token in redis: %w", op, err)
	}

	if bannedRefresh.ExpTime >= claims.ExpiresAt.Unix() {
		return "", "", fmt.Errorf("%s: %w", op, ErrTokenInBlackList)
	}

	version, err := s.versionProvider.RefreshVersion(ctx, claims.UID)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to find refresh token version: %w", op, HandleMssqlError(err))
	}

	if version != claims.Version {
		return "", "", fmt.Errorf("%s: refresh token version is not not equal to db: %w", op, ErrInvalidRefreshVersion)
	}

	access, refresh, err := s.generateTokens(ctx, claims.UID, claims.Version+1)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to generate tokens: %w", op, err)
	}

	return access, refresh, nil
}

func (s *Service) Tokens(
	ctx context.Context,
	uid int64,
) (string, string, error) {
	const op = "services.pcClub.auth.Tokens"

	version, err := s.versionProvider.RefreshVersion(ctx, uid)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to find refresh token version: %w", op, HandleMssqlError(err))
	}

	access, refresh, err := s.generateTokens(ctx, uid, version+1)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to generate tokens: %w", op, err)
	}

	return access, refresh, nil
}

func (s *Service) BanTokens(
	ctx context.Context,
	accessToken string,
	refreshToken string,
) (int64, error) {
	const op = "services.pcClub.auth.Tokens"

	claims, err := jwt.ParseToken(refreshToken, s.cfg.Refresh.Secret)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to parse token: %w", op, HandleMssqlError(err))
	}

	refresh := Refresh{
		ExpTime: claims.ExpiresAt.Unix(),
	}
	err = s.redisOwner.SetWithTTL(
		ctx,
		s.cfg.Refresh.TTL,
		&refresh,
		claims.UID,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to push refresh in black list: %w", op, err)
	}

	claims, err = jwt.ParseToken(accessToken, s.cfg.Access.Secret)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to parse token: %w", op, TokenError(err))
	}

	access := Access{
		ExpTime: claims.ExpiresAt.Unix(),
	}
	err = s.redisOwner.SetWithTTL(
		ctx,
		s.cfg.Access.TTL,
		&access,
		claims.UID,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to push access in black list: %w", op, err)
	}

	return claims.UID, nil
}
