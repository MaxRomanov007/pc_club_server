package mssql

import (
	"context"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/api/database/gorm"
)

func (s *Storage) Dishes(
	ctx context.Context,
	limit int,
	offset int,
) ([]models.Dish, error) {
	const op = "storage.mssql.Dishes"

	db := s.db.WithContext(ctx).Preload("DishImages")
	if limit > 0 {
		db = db.Limit(limit)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}
	var dishes []models.Dish
	if res := db.Find(&dishes); gorm.IsFailResult(res) {
		return nil, fmt.Errorf("%s: failed to get dishes: %w", op, errorByResult(res))
	}

	for i := 0; i < len(dishes); i++ {
		for j := 0; j < len(dishes[i].DishImages); j++ {
			dishes[i].DishImages[j].Path = s.cfg.Images.Host + "/" + dishes[i].DishImages[j].Path
		}
	}

	return dishes, nil
}
