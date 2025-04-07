package order

import (
	"context"
	"pc_club_server/internal/config"
	"pc_club_server/internal/domain/models"
)

type PcOwner interface {
	OrderPc(
		ctx context.Context,
		order *models.PcOrder,
	) (err error)
}

type PcProvider interface {
	PcHourCost(
		ctx context.Context,
		pcId int64,
	) (cost float32)
}

type DishOwner interface {
	OrderDish(
		ctx context.Context,
		order *models.DishOrder,
		orderList *models.DishOrderList,
	) (err error)
}

type DishProvider interface {
	DishCost(
		ctx context.Context,
		id int64,
	) (cost float32)
}

type Service struct {
	cfg          *config.OrdersConfig
	pcOwner      PcOwner
	pcProvider   PcProvider
	dishOwner    DishOwner
	dishProvider DishProvider
}

func NewService(
	cfg *config.OrdersConfig,
	pcOwner PcOwner,
	pcProvider PcProvider,
	dishOwner DishOwner,
	dishProvider DishProvider,
) *Service {
	return &Service{
		cfg:          cfg,
		pcOwner:      pcOwner,
		pcProvider:   pcProvider,
		dishOwner:    dishOwner,
		dishProvider: dishProvider,
	}
}
