package config

import (
	"os"
)

type Config struct {
	Origin string
}

func FromEnv() Config {
	return Config{
		Origin: env("Origin", "localhost:4200"),
	}
}

func env(name, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	return value
}
