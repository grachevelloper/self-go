package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App      AppConfig
	HTTP     HTTPConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	Env string
}

type HTTPConfig struct {
	Port   int
	Origin string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() (*Config, error) {
	httpPort, err := getInt("HTTP_PORT", 8080)
	if err != nil {
		return nil, err
	}

	postgresPort, err := getInt("POSTGRES_PORT", 5433)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		App: AppConfig{
			Env: getString("APP_ENV", "development"),
		},
		HTTP: HTTPConfig{
			Port:   httpPort,
			Origin: getString("ORIGIN", "localhost:4200"),
		},
		Postgres: PostgresConfig{
			Host:     getString("POSTGRES_HOST", "localhost"),
			Port:     postgresPort,
			User:     getString("POSTGRES_USER", "postgres"),
			Password: getString("POSTGRES_PASSWORD", "postgres"),
			DBName:   getString("POSTGRES_DB", "postgres"),
			SSLMode:  getString("POSTGRES_SSL_MODE", "disable"),
		},
	}

	return cfg, nil
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
	)
}

func getString(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid env %s: %w", key, err)
	}

	return parsed, nil
}
