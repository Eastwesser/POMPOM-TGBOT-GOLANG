package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Инициализация соединения с PostgreSQL

func ConnectDB(cfg *Config) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Connected to the database successfully!")
	return dbpool
}
