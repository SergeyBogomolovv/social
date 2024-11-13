package app

import (
	"github.com/jmoiron/sqlx"
)

type App struct {
	addr string
	db   *sqlx.DB
}

func NewApp(addr string, db *sqlx.DB) *App {
	return &App{
		addr: addr,
		db:   db,
	}
}
