package config

import (
	"os"
)

type Config struct {
	Environment    string
	LogLevel       string
	KafkaBroker    string
	MinIOEndpoint  string
	MinIOAccessKey string
	MinIOSecretKey string
}

func Load() (*Config, error) {
	return &Config{
		Environment:    getEnv("ENVIRONMENT", "development"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		KafkaBroker:    getEnv("KAFKA_BROKER", "localhost:9092"),
		MinIOEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIOAccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MinIOSecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
