package app

import (
	"net/http"
	"social/internal/http/controller"
	"social/internal/http/usecase"
	proto "social/pkg/proto/generated"
)

type App struct {
	addr        string
	usersClient proto.UserServiceClient
	postsClient proto.PostServiceClient
}

func (a *App) Run() error {
	router := http.NewServeMux()

	usersUsecase := usecase.NewUsersUsecase(a.usersClient)
	controller.RegisterUsersController(router, usersUsecase)

	server := &http.Server{
		Addr:    a.addr,
		Handler: router,
	}

	return server.ListenAndServe()
}

func NewApp(addr string, usersClient proto.UserServiceClient, postsClient proto.PostServiceClient) *App {
	return &App{
		addr:        addr,
		usersClient: usersClient,
		postsClient: postsClient,
	}
}
