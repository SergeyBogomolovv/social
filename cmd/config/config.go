package config

import (
	"social/pkg/utils"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURI string
	PostsPort   string
	UsersPort   string
	HttpPort    string
	JwtSecret   string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		PostgresURI: utils.GetEnvString("POSTGRES_URI", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		PostsPort:   utils.GetEnvString("POSTS_PORT", ":9001"),
		UsersPort:   utils.GetEnvString("USERS_PORT", ":9000"),
		HttpPort:    utils.GetEnvString("HTTP_PORT", ":8080"),
		JwtSecret:   utils.GetEnvString("JWT_SECRET", "secret"),
	}
}
