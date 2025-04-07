package mssql

import (
	"context"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/api/database/gorm"
	gorm2 "pc_club_server/internal/lib/api/database/gorm"
)

func (s *Storage) OrderPc(
	ctx context.Context,
	order *models.PcOrder,
) (err error) {
	const op = "storage.mssql.OrderPc"

	tx := s.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		commitErr := tx.Commit().Error
		if commitErr != nil {
			err = fmt.Errorf("%s: %w", op, commitErr)
		}
	}()

	if res := tx.WithContext(ctx).
		Exec(`
			UPDATE dbo.users
			SET balance = balance - ? 
			WHERE user_id = ? AND balance >= ?`,
			order.Cost, order.UserID, order.Cost,
		); gorm2.IsFailResult(res) {

		if res.RowsAffected == 0 {
			return fmt.Errorf("%s: %w", op, ErrCheckConstraintViolated)
		}

		return fmt.Errorf("%s: failed to update balance: %w", op, errorByResult(res))
	}

	if res := tx.WithContext(ctx).Save(order); gorm.IsFailResult(res) {
		return fmt.Errorf("%s: %w", op, errorByResult(res))
	}

	return nil
}
