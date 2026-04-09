package config

import (
	"os"
	"strconv"
)

type Config struct {
	Environment      string
	Port             string
	DBConfig         *DBConfig
	MigrationEnabled bool
	MigrationPath    string
}

func (c Config) IsEnvLocal() bool {
	return c.Environment == "local"
}

type DBConfig struct {
	User     string
	Name     string
	Host     string
	Port     int64
	Password string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		User:     GetRequiredEnv("DB_USER"),
		Name:     GetRequiredEnv("DB_NAME"),
		Port:     GetIntEnvWithDefault("DB_PORT", 5432),
		Password: GetRequiredEnv("DB_PASSWORD"),
		Host:     GetRequiredEnv("DB_HOST"),
	}
}

func Load() *Config {
	return &Config{
		Environment:      GetRequiredEnv("APP_ENV"),
		Port:             GetEnvWithDefault("APP_PORT", "8080"),
		DBConfig:         LoadDBConfig(),
		MigrationEnabled: GetBoolEnvWithDefault("MIGRATION_ENABLED", false),
		MigrationPath:    GetEnvWithDefault("MIGRATION_PATH", "./internal/database/migration"),
	}
}

func GetEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetRequiredEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	panic("missing required environment variable: " + key)
}

func GetIntEnvWithDefault(key string, defaultValue int64) int64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return int64(intVal)
}

func GetBoolEnvWithDefault(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return boolVal
}
