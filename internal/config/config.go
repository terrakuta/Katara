package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGO_URI      string
	REDIS_URI      string
	PORT           string
	REDIS_PASSWORD string
	REDIS_DB       int
	MONGO_DB       string
	SYNC_ENABLED   bool
}

func getEnvBool(key string, def bool) bool {
	a := os.Getenv(key)

	if a == "" {
		return def
	}

	b, err := strconv.ParseBool(a)
	if err != nil {
		return def
	}
	return b
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading enf files")
	}
	return &Config{
		MONGO_URI:      os.Getenv("MONGO_URI"),
		PORT:           os.Getenv("PORT"),
		MONGO_DB:       os.Getenv("MONGO_DB"),
		REDIS_URI:      os.Getenv("REDIS_URI"),
		REDIS_PASSWORD: os.Getenv("REDIS_PASSWORD"),
		REDIS_DB:       0,
		SYNC_ENABLED:   getEnvBool("SYNC_ENABLED", false),
	}, nil
}
