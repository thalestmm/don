package main

import (
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// TODO: Pick up from env
const (
	CURRENCY            string = "USD"
	FUCK_IT_MONEY_CENTS int    = 200000000
)

type Config struct {
	Currency string `env:"CURRENCY" envDefault:"USD"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("Error reading .env file")
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		slog.Error(err.Error())
		return nil, fmt.Errorf("Error parsing environment into Config")
	}
	return &cfg, nil
}
