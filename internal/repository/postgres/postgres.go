package postgres

import (
	"database/sql"
	"fmt"

	"github.com/himmel520/wb_L0/internal/config"
)

func NewPostgres(cfg config.Postgres) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%v dbname=%v sslmode=%v user=%v password=%v",
			cfg.Host, cfg.DBName, cfg.SSLMode, cfg.User, cfg.Password))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
