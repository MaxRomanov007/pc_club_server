package mssql

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pc_club_server/internal/config"
	"pc_club_server/internal/lib/api/database/mssql"
)

type Storage struct {
	cfg *config.Config
	db  *gorm.DB
}

func New(cfg *config.Config) (*Storage, error) {
	const op = "storage.mssql.New"

	db, err := gorm.Open(sqlserver.Open(mssql.GenerateConnString(cfg.Database.SQLServer)), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: failed to connect: %w", op, err)
	}

	return &Storage{
		db:  db,
		cfg: cfg,
	}, nil
}
