package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(uri string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("error ping database: %v", err)
	}

	return conn, nil
}
