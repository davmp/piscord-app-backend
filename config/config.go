package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	MongoURI string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		Port:     getEnv("PORT", ""),
		MongoURI: getEnv("MONGO_URI", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
