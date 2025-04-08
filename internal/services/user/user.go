package user

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/storage/mssql"
)

func (s *Service) User(
	ctx context.Context,
	uid int64,
) (models.User, error) {
	const op = "services.user.User"

	user, err := s.userProvider.User(ctx, uid)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: failed to get user from mssql: %w", op, HandleStorageError(err))
	}

	return user, nil
}

func (s *Service) UserByEmail(
	ctx context.Context,
	email string,
) (models.User, error) {
	const op = "services.user.UserByEmail"

	user, err := s.userProvider.UserByEmail(ctx, email)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: failed to get user by email from mssql: %w", op, HandleStorageError(err))
	}

	return user, nil
}

func (s *Service) SaveUser(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "services.user.SaveUser"

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("%s: failed hash password: %w", op, err)
	}

	id, err := s.userOwner.SaveUser(
		ctx,
		&models.User{
			Email:    email,
			Password: passHash,
		})
	if err != nil {
		return 0, fmt.Errorf("%s: failed to save user in mssql: %w", op, HandleStorageError(err))
	}

	return id, nil
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "services.user.Login"

	user, err := s.userProvider.UserByEmail(ctx, email)
	if errors.Is(err, mssql.ErrNotFound) {
		return 0, fmt.Errorf("%s: failed to get user from mssql: %w", op, ErrInvalidCredentials)
	}
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get user from mssql: %w", op, HandleStorageError(err))
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return 0, fmt.Errorf("%s: failed to compare password: %w", op, HandleStorageError(err))
	}

	return user.UserID, nil
}

func (s *Service) DeleteUser(
	ctx context.Context,
	uid int64,
) error {
	const op = "services.user.DeleteUser"

	err := s.userOwner.DeleteUser(ctx, uid)
	if err != nil {
		return fmt.Errorf("%s: failed to delete user from mssql: %w", op, HandleStorageError(err))
	}

	return nil
}

func (s *Service) IsAdmin(
	ctx context.Context,
	uid int64,
) error {
	const op = "services.user.IsAdmin"

	role, err := s.userProvider.UserRole(ctx, uid)
	if err != nil {
		return fmt.Errorf("%s: failed to get user role from mssql: %w", op, HandleStorageError(err))
	}

	if role != AdminRoleName {
		return fmt.Errorf("%s: user is not admin: %w", op, ErrAccessDenied)
	}

	return nil
}

func (s *Service) UserWithOrders(
	ctx context.Context,
	uid int64,
) (models.User, error) {
	const op = "services.user.User"

	user, err := s.userProvider.UserWithOrders(ctx, uid)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: failed to get user from mssql: %w", op, HandleStorageError(err))
	}

	return user, nil
}

func (s *Service) AddMoney(
	ctx context.Context,
	uid int64,
	count float32,
) error {
	const op = "services.user.AddMoney"

	if err := s.userOwner.AddUserMoney(ctx, uid, count); err != nil {
		return fmt.Errorf("%s: %w", op, HandleStorageError(err))
	}

	return nil
}
