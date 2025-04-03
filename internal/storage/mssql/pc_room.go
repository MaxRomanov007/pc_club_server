package mssql

import (
	"context"
	"fmt"
	"pc_club_server/internal/domain/models"
)

func (s *Storage) PcRooms(
	ctx context.Context,
	pcTypeId int64,
) ([]models.PcRoom, error) {
	const op = "storage.mssql.pc_room.PcRooms"

	var rooms []models.PcRoom

	err := s.db.WithContext(ctx).
		Joins("JOIN pc ON pc.pc_room_id = pc_rooms.pc_room_id").
		Joins("JOIN pc_types ON pc_types.pc_type_id = pc.pc_type_id").
		Where("pc.pc_type_id = ?", pcTypeId).
		Preload("Pcs").
		Find(&rooms).Error

	if err != nil {
		return nil, fmt.Errorf("%s: failed to get pc rooms: %w", op, err)
	}

	return rooms, nil
}
