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

type Service struct {
	cfg        *config.OrdersConfig
	pcOwner    PcOwner
	pcProvider PcProvider
}

func NewService(
	cfg *config.OrdersConfig,
	pcOwner PcOwner,
	pcProvider PcProvider,
) *Service {
	return &Service{
		cfg:        cfg,
		pcOwner:    pcOwner,
		pcProvider: pcProvider,
	}
}
