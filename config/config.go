package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort   string
	JikanAPIURL  string
	CacheTimeout int
}

func Load() *Config {
	return &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		JikanAPIURL:  getEnv("JIKAN_API_URL", "https://api.jikan.moe/v4"),
		CacheTimeout: getEnvInt("CACHE_TIMEOUT_MINUTES", 10),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
