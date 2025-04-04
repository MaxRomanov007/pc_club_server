package user

import (
	"context"
	"pc_club_server/internal/domain/models"
)

type provider interface {
	User(
		ctx context.Context,
		uid int64,
	) (user models.User, err error)

	UserByEmail(
		ctx context.Context,
		email string,
	) (user models.User, err error)

	UserRole(
		ctx context.Context,
		uid int64,
	) (role string, err error)
}

type owner interface {
	SaveUser(
		ctx context.Context,
		user *models.User,
	) (id int64, err error)

	DeleteUser(
		ctx context.Context,
		uid int64,
	) (err error)
}

type Service struct {
	userProvider provider
	userOwner    owner
}

const (
	AdminRoleName = "admin"
)

func NewService(
	userProvider provider,
	userOwner owner,
) *Service {
	return &Service{
		userProvider: userProvider,
		userOwner:    userOwner,
	}
}
