package main

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"pc_club_server/internal/config"
	"pc_club_server/internal/lib/api/database/mssql"
)

func main() {
	cfg := config.MustLoad()

	migPath := os.Getenv("MIGRATIONS_PATH")
	if migPath == "" {
		log.Fatal("migrations path not set")
	}

	m, err := migrate.New(
		"file://"+migPath,
		mssql.GenerateConnString(cfg.Database.SQLServer),
	)
	if err != nil {
		log.Fatal("failed to create migrator: " + err.Error())
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("db is up to date")
			return
		}
		log.Fatal("failed to run migrations: " + err.Error())
	}
	log.Println("db successfully migrated")
}
