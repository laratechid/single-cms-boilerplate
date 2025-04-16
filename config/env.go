package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port string
		Env  string
		Name string
	}
	Postgres struct {
		Host string
		Port string
		DB   string
		User string
		Pass string
	}
	Jwt struct {
		Secret string
	}
}

func Env() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{}

	cfg.App.Port = os.Getenv("APP_PORT")
	cfg.App.Env = os.Getenv("APP_ENV")
	cfg.App.Name = os.Getenv("APP_NAME")

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.DB = os.Getenv("POSTGRES_DB")
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Pass = os.Getenv("POSTGRES_PASS")

	cfg.Jwt.Secret = os.Getenv("JWT_SECRET")

	return cfg
}
