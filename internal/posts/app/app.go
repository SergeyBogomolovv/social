package app

import (
	"log/slog"
	"net"
	"social/internal/posts/controller"
	"social/internal/posts/repository"
	"social/internal/posts/usecase"

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
	repo := repository.NewPostsRepository(a.db)
	usecase := usecase.NewPostsUsecase(repo)
	controller.RegisterPostsController(a.grpc, usecase)

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
