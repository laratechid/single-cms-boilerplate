package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port       string
		Env        string
		Name       string
		ApiKey     string
		TempoToken string
	}
	Postgres struct {
		Host string
		Port string
		DB   string
		User string
		Pass string
	}

	Redis struct {
		Host string
		Port string
		User string
		Pass string
	}
	Jwt struct {
		Secret string
	}
	BaseUrl struct {
		FrontendUrl string
	}
	Csrf struct {
		TokenDuration string
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
	cfg.App.ApiKey = os.Getenv("APP_API_KEY")
	cfg.App.TempoToken = os.Getenv("APP_TEMPO_TOKEN")

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.DB = os.Getenv("POSTGRES_DB")
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Pass = os.Getenv("POSTGRES_PASS")

	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Port = os.Getenv("REDIS_PORT")
	cfg.Redis.User = os.Getenv("REDIS_USER")
	cfg.Redis.Pass = os.Getenv("REDIS_PASS")

	cfg.Jwt.Secret = os.Getenv("JWT_SECRET")

	cfg.BaseUrl.FrontendUrl = os.Getenv("FRONTEND_URL")

	cfg.Csrf.TokenDuration = os.Getenv("CSRF_TOKEN_DURATION")

	return cfg
}
