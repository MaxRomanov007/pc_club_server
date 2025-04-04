package order

import (
	"context"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/lib/codes"
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
		UserID:        uid,
		PcID:          pcId,
		Code:          code,
		Cost:          s.pcProvider.PcHourCost(ctx, pcId),
		StartTime:     time.Now(),
		Duration:      HourCount,
		ActualEndTime: time.Now().Add(time.Duration(HourCount) * time.Hour),
		OrderDate:     time.Now(),
	}

	if err := s.pcOwner.OrderPc(ctx, order); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return code, nil
}
