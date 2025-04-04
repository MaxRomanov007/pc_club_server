package mssql

import (
	"context"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/api/database/gorm"
)

func (s *Storage) OrderPc(
	ctx context.Context,
	order *models.PcOrder,
) error {
	const op = "storage.mssql.OrderPc"

	if res := s.db.WithContext(ctx).Save(order); gorm.IsFailResult(res) {
		return fmt.Errorf("%s: %w", op, errorByResult(res))
	}

	return nil
}
