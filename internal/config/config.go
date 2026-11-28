package config

import (
	"errors"
	"os"
)

type Config struct {
	HTTPPort    string
	DatabaseURL string
	LogLevel    string
}

func Load() (Config, error) {
	cfg := Config{
		HTTPPort:    os.Getenv("HTTP_PORT"),
		DatabaseURL: os.Getenv("DB_URL"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
	}
	if cfg.HTTPPort == "" {
		cfg.HTTPPort = ":4444"
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}
	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("DB_URL is required")
	}
	return cfg, nil
}
