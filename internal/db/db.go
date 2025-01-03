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

func createDBConnection(cfg config.DBConfig) (*sql.DB, error) {
	log.Println("Try connect to database...")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	return sql.Open("postgres", connStr)
}

func NewRepository(cfg config.DBConfig) (*Repository, error) {
	db, err := createDBConnection(cfg)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("successfully connected to database '%s'\n", cfg.Name)
	return &Repository{db: db}, nil
}
