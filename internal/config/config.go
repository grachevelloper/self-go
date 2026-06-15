package config

import (
	"os"
)

type Config struct {
	TokenSecret string
}

func FromEnv() Config {
	return Config{
		TokenSecret: env("TOKEN_SECRET", "123"),
	}
}

func env(name, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	return value
}
