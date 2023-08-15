package config

import (
	"fmt"
	"os"
)

var (
	postgresHost     = os.Getenv("POSTGRES_HOST")
	postgresUser     = os.Getenv("POSTGRES_USER")
	postgresPassword = os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase = os.Getenv("POSTGRES_DATABASE")
	postgresPort     = os.Getenv("POSTGRES_PORT")
	PostgesDSN       = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort,
	)
)
