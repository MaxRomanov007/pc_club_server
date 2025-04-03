package mssql

import (
	"context"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/api/database/gorm"
)

func (s *Storage) PcTypes(
	ctx context.Context,
	limit int,
	offset int,
) ([]models.PcType, error) {
	const op = "storage.mssql.pc_type.PcTypes"

	db := s.db.WithContext(ctx).Preload("PcTypeImages")
	if limit > 0 {
		db = db.Limit(limit)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}
	var types []models.PcType
	if res := db.Find(&types); gorm.IsFailResult(res) {
		return nil, fmt.Errorf("%s: failed to get pc types: %w", op, errorByResult(res))
	}

	for i := 0; i < len(types); i++ {
		for j := 0; j < len(types[i].PcTypeImages); j++ {
			types[i].PcTypeImages[j].Path = s.cfg.Images.Host + "/" + types[i].PcTypeImages[j].Path
		}
	}

	return types, nil
}

func (s *Storage) PcType(
	ctx context.Context,
	typeID int64,
) (models.PcType, error) {
	const op = "storage.mssql.pc_type.PcType"

	var pcType models.PcType
	if res := s.db.WithContext(ctx).Preload("PcTypeImages").
		Preload("Processor").Preload("Processor.ProcessorProducer").
		Preload("VideoCard").Preload("VideoCard.VideoCardProducer").
		Preload("Monitor").Preload("Monitor.MonitorProducer").
		Preload("RAM").Preload("RAM.RAMType").
		First(&pcType, typeID); gorm.IsFailResult(res) {

		return models.PcType{}, fmt.Errorf("%s: failed to get pc type: %w", op, errorByResult(res))
	}

	for j := 0; j < len(pcType.PcTypeImages); j++ {
		pcType.PcTypeImages[j].Path = s.cfg.Images.Host + "/" + pcType.PcTypeImages[j].Path
	}

	return pcType, nil
}
