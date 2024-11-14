package main

import (
	"log/slog"
	"social/cmd/config"
	"social/internal/database"
	"social/internal/users/app"
)

func main() {
	cfg := config.NewConfig()

	db, err := database.NewPostgresConnection(cfg.PostgresURI)
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		panic(err)
	}

	app := app.NewApp(cfg.UsersPort, db)

	err = app.Run()
	if err != nil {
		slog.Error("error running app", "error", err)
		panic(err)
	}
}
