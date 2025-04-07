package order

import (
	"context"
	"errors"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/codes"
	"pc_club_server/internal/storage/mssql"
	"time"
)

func (s *Service) OrderPc(
	ctx context.Context,
	uid int64,
	pcId int64,
	HourCount int16,
) (string, error) {
	const op = "services.order.OrderPc"

	code := codes.Generate(s.cfg.CodeLength)

	order := &models.PcOrder{
		UserID:          uid,
		PcID:            pcId,
		PcOrderStatusID: 1,
		Code:            code,
		Cost:            s.pcProvider.PcHourCost(ctx, pcId) * float32(HourCount),
		StartTime:       time.Now(),
		Duration:        HourCount,
		EndTime:         time.Now().Add(time.Duration(HourCount) * time.Hour),
		ActualEndTime:   time.Now().Add(time.Duration(HourCount) * time.Hour),
		OrderDate:       time.Now(),
	}

	if err := s.pcOwner.OrderPc(ctx, order); err != nil {
		if errors.Is(err, mssql.ErrCheckConstraintViolated) {
			return "", fmt.Errorf("%s: %w", op, ErrNotEnoughMoney)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return code, nil
}
