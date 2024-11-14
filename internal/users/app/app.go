package app

import (
	"log/slog"
	"net"
	"social/internal/users/controller"
	"social/internal/users/repository"
	"social/internal/users/usecase"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type App struct {
	addr string
	db   *sqlx.DB
	grpc *grpc.Server
}

func (a *App) Run() error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		slog.Error("failed to listen", "error", err)
		return err
	}

	repo := repository.NewUsersRepository(a.db)
	usecase := usecase.NewUsersUsecase(repo)
	controller.RegisterUsersController(a.grpc, usecase)

	slog.Info("Starting gRPC server", "address", a.addr)
	return a.grpc.Serve(listener)
}

func NewApp(addr string, db *sqlx.DB) *App {
	return &App{
		addr: addr,
		db:   db,
		grpc: grpc.NewServer(),
	}
}
