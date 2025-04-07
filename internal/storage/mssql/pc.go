package mssql

import (
	"context"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/api/database/gorm"
)

func (s *Storage) PcHourCost(
	ctx context.Context,
	pcId int64,
) float32 {
	var pc models.Pc
	if res := s.db.WithContext(ctx).
		Preload("PcType").
		First(&pc, pcId); gorm.IsFailResult(res) {

		return 0
	}

	return pc.PcType.HourCost
}

func (s *Storage) Pcs(
	ctx context.Context,
	typeId int64,
) ([]models.Pc, error) {
	const op = "storage.mssql.Pcs"

	var pcs []models.Pc
	if res := s.db.WithContext(ctx).
		Where("pc_type_id = ?", typeId).
		Preload("PcRoom").
		Find(&pcs); gorm.IsFailResult(res) {

		return nil, fmt.Errorf("%s: %w", op, errorByResult(res))
	}

	return pcs, nil
}
