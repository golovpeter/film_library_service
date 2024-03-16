package common

import (
	"fmt"

	"github.com/golovpeter/vk_intership_test_task/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

func CreateDbClient(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database))

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err = goose.Up(db.DB, "migrations"); err != nil {
		return nil, err
	}

	return db, nil
}
