package app

import (
	"context"
	"fmt"
	"log/slog"
	"pc_club_server/internal/app/pcClub"
	"pc_club_server/internal/config"
	pcClubServer "pc_club_server/internal/http-server/handlers"
	"pc_club_server/internal/services/auth"
	"pc_club_server/internal/services/order"
	"pc_club_server/internal/services/user"
	gorm "pc_club_server/internal/storage/mssql"
	"pc_club_server/internal/storage/redis"
)

type App struct {
	PCClub *pcClub.App
}

func MustLoad(
	ctx context.Context,
	log *slog.Logger,
	cfg *config.Config,
) *App {
	app, err := New(ctx, log, cfg)
	if err != nil {
		panic(err)
	}

	return app
}

func New(
	ctx context.Context,
	log *slog.Logger,
	cfg *config.Config,
) (*App, error) {
	const op = "app.New"

	mssqlStorage, err := gorm.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create mssql storage: %w", op, err)
	}

	redisStorage, err := redis.New(ctx, cfg.Database.Redis)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create redis storage: %w", op, err)
	}

	authService := auth.NewService(cfg.Auth, redisStorage, redisStorage, mssqlStorage, mssqlStorage)
	userService := user.NewService(mssqlStorage, mssqlStorage)
	orderService := order.NewService(cfg.Orders, mssqlStorage, mssqlStorage, mssqlStorage, mssqlStorage)

	pcClubApi := pcClubServer.New(
		log,
		cfg,
		userService,
		authService,
		mssqlStorage,
		mssqlStorage,
		orderService,
		mssqlStorage,
		orderService,
	)

	pcClubApplication := pcClub.New(cfg.HttpsServer, pcClubApi)

	return &App{
		PCClub: pcClubApplication,
	}, nil
}
