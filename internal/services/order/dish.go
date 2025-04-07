package order

import (
	"context"
	"errors"
	"fmt"
	"pc_club_server/internal/domain/models"
	"pc_club_server/internal/storage/mssql"
	"time"
)

func (s *Service) OrderDish(
	ctx context.Context,
	uid int64,
	dishId int64,
	count int16,
) error {
	const op = "services.order.OrderDish"

	order := models.DishOrder{
		DishOrderStatusID: 1,
		UserID:            uid,
		Cost:              s.dishProvider.DishCost(ctx, dishId),
		OrderDate:         time.Now(),
	}

	orderList := models.DishOrderList{
		DishID: dishId,
		Count:  count,
	}

	if err := s.dishOwner.OrderDish(ctx, &order, &orderList); err != nil {
		if errors.Is(err, mssql.ErrCheckConstraintViolated) {
			return fmt.Errorf("%s: %w", op, ErrNotEnoughMoney)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
