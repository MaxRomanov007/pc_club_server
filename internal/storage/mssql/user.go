package mssql

import (
	"context"
	"fmt"
	gorm2 "gorm.io/gorm"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/api/database/gorm"
)

func (s *Storage) SaveUser(
	ctx context.Context,
	user *models.User,
) (int64, error) {
	const op = "storage.user.mssql.SaveUser"

	if res := s.db.WithContext(ctx).Save(user); gorm.IsFailResult(res) {
		return 0, fmt.Errorf("%s: failed to save user: %w", op, errorByResult(res))
	}

	return user.UserID, nil
}

func (s *Storage) User(
	ctx context.Context,
	uid int64,
) (models.User, error) {
	const op = "storage.mssql.user.User"

	var user models.User
	if res := s.db.WithContext(ctx).First(&user, uid); gorm.IsFailResult(res) {
		return models.User{}, fmt.Errorf("%s: failed to get user: %w", op, errorByResult(res))
	}

	return user, nil
}

func (s *Storage) UserRole(
	ctx context.Context,
	uid int64,
) (string, error) {
	const op = "storage.mssql.user.UserRole"

	var user models.User
	if res := s.db.WithContext(ctx).
		Preload("UserRole").
		First(&user, uid); gorm.IsFailResult(res) {

		return "", fmt.Errorf("%s: failed to get users role: %w", op, errorByResult(res))
	}

	return user.UserRole.Name, nil
}

func (s *Storage) RefreshVersion(
	ctx context.Context,
	uid int64,
) (int64, error) {
	const op = "storage.mssql.user.RefreshVersion"

	var version int64
	if res := s.db.WithContext(ctx).
		Model(&models.User{}).
		Select("refresh_token_version").
		First(&version, uid); gorm.IsFailResult(res) {

		return 0, fmt.Errorf("%s: failed to get users refresh token version: %w", op, errorByResult(res))
	}

	return version, nil
}

func (s *Storage) UserByEmail(
	ctx context.Context,
	email string,
) (models.User, error) {
	const op = "storage.mssql.user.UserByEmail"

	var user models.User
	if res := s.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user); gorm.IsFailResult(res) {

		return models.User{}, fmt.Errorf("%s: failed to get user by email: %w", op, errorByResult(res))
	}

	return user, nil
}

func (s *Storage) UpdateEmail(
	ctx context.Context,
	uid int64,
	email string,
) error {
	const op = "storage.mssql.user.UpdateEmail"

	if res := s.db.WithContext(ctx).
		Model(models.User{}).
		Where("user_id = ?", uid).
		UpdateColumn("email", email); gorm.IsFailResult(res) {

		return fmt.Errorf("%s: failed to update email: %w", op, errorByResult(res))
	}

	return nil
}

func (s *Storage) IncRefreshVersion(
	ctx context.Context,
	uid int64,
) error {
	const op = "storage.mssql.user.IncRefreshVersion"

	sql := "UPDATE dbo.users SET refresh_token_version = refresh_token_version + 1 WHERE user_id = ?"
	if res := s.db.WithContext(ctx).Exec(sql, uid); gorm.IsFailResult(res) {
		return fmt.Errorf("%s: failed to get user by email: %w", op, errorByResult(res))
	}

	return nil
}

func (s *Storage) DeleteUser(
	ctx context.Context,
	uid int64,
) error {
	const op = "storage.mssql.user.DeleteUser"

	if res := s.db.WithContext(ctx).Delete(&models.User{}, uid); gorm.IsFailResult(res) {
		return fmt.Errorf("%s: failed to delete user: %w", op, errorByResult(res))
	}

	return nil
}

func (s *Storage) UserWithOrders(
	ctx context.Context,
	uid int64,
) (models.User, error) {
	const op = "storage.mssql.user.User"

	var user models.User
	if res := s.db.WithContext(ctx).
		Preload("PcOrders").
		Preload("PcOrders.Pc").
		Preload("PcOrders.Pc.PcType").
		Preload("DishOrders").
		Preload("DishOrders.DishOrderList").
		Preload("DishOrders.DishOrderList.Dish").
		First(&user, uid); gorm.IsFailResult(res) {
		return models.User{}, fmt.Errorf("%s: failed to get user: %w", op, errorByResult(res))
	}

	return user, nil
}

func (s *Storage) AddUserMoney(
	ctx context.Context,
	uid int64,
	count float32,
) error {
	const op = "storage.mssql.AddUserMoney"

	if res := s.db.WithContext(ctx).First(&models.User{}, uid).
		Update("balance", gorm2.Expr("balance + ?", count)); gorm.IsFailResult(res) {

		return fmt.Errorf("%s: %w", op, errorByResult(res))
	}

	return nil
}
