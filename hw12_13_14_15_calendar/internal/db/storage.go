package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/config"
)

func PostgreSQLConnectFromConfig(ctx context.Context, configFile string) (*sql.DB, error) {
	config, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	if config.Storage.DriverName == "postgresql" {
		return PostgreSQLConnect(ctx, config.Storage.Dsn)
	}
	return nil, errors.New("в конфиге выбрана база данных не postrgresql")
}

func PostgreSQLConnect(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to load driver: %w", err)
	}
	return db, nil
}

func Close(ctx context.Context, db *sql.DB) error {
	return db.Close()
}
