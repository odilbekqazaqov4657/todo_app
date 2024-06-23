package db

import (
	"context"
	"log"

	"github.com/odilbekqazaqov4657/todo_app/config"

	"fmt"

	"github.com/jackc/pgx"
)

func NewPgx(cfg config.Config) (*pgx.Conn, error) {

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DbUser,
		cfg.DbUserPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	con, err := pgx.Connect(context.Background(), url)

	if err != nil {
		log.Println("error on connecting to db !", err)

		return nil, err

	}
	return con, nil
}
