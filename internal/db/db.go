package db

import (
	"database/sql"
	"fmt"
	"log"
	"project_sem/internal/config"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg config.DBConfig) (*Repository, error) {
	log.Println("Connecting to database...")

	connStr := buildConnectionString(cfg)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Printf("Successfully connected to database '%s'\n", cfg.Name)
	return &Repository{db: db}, nil
}

func buildConnectionString(cfg config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)
}
