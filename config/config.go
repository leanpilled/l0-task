package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DATABASE_URL string

	HOST string `env:"HOST"`
	PORT string `env:"PORT"`

	ClusterID string `env:"CLUSTERID"`
	ClientID  string `env:"CLIENTID"`
}

func New() Config {
	ctx := context.Background()
	envErr := godotenv.Load(".env")
	if envErr != nil {
		panic("Error loading .env file")
	}

	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(err)
	}
	cfg.DATABASE_URL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("POSTGRES_DB"))
	return cfg
}
