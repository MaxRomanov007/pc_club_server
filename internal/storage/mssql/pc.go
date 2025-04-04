package mssql

import (
	"context"
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
