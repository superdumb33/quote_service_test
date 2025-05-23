package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//add field validation
type AppCfg struct {
	AppPort          int
	PostgresUser     string
	PostgresDB       string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
}

// it'll throw a panic if something goes wrong
func MustInit() AppCfg {
	//.Load() should be called if the app is being launched with `go run`; docker compose will launch service with env variables set from provided .env file
	godotenv.Load(".env")
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

	return AppCfg{
		AppPort: appPort,
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresDB: os.Getenv("POSTGRES_DB"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
	}
}