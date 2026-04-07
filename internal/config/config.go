package config

import (
	"os"
	"strconv"
)

type Config struct {
	Environment string
	Port        string
	DBConfig    *DBConfig
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
		User:     GetEnvWithDefault("DB_USER", "postgres"),
		Name:     GetEnvWithDefault("DB_NAME", "student_main"),
		Port:     GetIntEnvWithDefault("DB_PORT", 5432),
		Password: GetEnvWithDefault("DB_PASSWORD", ""),
		Host:     GetEnvWithDefault("DB_HOST", "localhost"),
	}
}

func Load() *Config {
	return &Config{
		Environment: GetEnvWithDefault("APP_ENV", "development"),
		Port:        GetEnvWithDefault("APP_PORT", "8080"),
		DBConfig:    LoadDBConfig(),
	}
}

func GetEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
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
