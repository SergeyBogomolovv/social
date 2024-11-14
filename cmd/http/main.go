package main

import (
	"log/slog"
	"social/cmd/config"
	"social/internal/http/app"
	proto "social/pkg/proto/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.NewConfig()
	usersConn, err := grpc.NewClient(cfg.UsersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to connect to users service")
		panic(err)
	}
	defer usersConn.Close()
	usersClient := proto.NewUserServiceClient(usersConn)

	postsConn, err := grpc.NewClient(cfg.PostsPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to connect to posts service")
		panic(err)
	}
	defer postsConn.Close()
	postsClient := proto.NewPostServiceClient(postsConn)

	app := app.NewApp(cfg.HttpPort, usersClient, postsClient)

	slog.Info("Starting HTTP server", "address", cfg.HttpPort)

	err = app.Run()
	if err != nil {
		slog.Error("error running app", "error", err)
	}
}
