package config

import (
	"os"
)

var PostgresURL = os.Getenv("POSTGRES_URL")
