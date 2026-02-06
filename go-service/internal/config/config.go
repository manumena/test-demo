package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	ServicePort string

	BackendBaseURL string
	BackendEmail   string
	BackendPassword string

	BackendTimeout time.Duration
}

func Load() *Config {
	cfg := &Config{
		ServicePort: getEnv("GO_SERVICE_PORT", "8080"),

		BackendBaseURL:  mustGetEnv("BACKEND_BASE_URL"),
		BackendEmail:    mustGetEnv("BACKEND_EMAIL"),
		BackendPassword: mustGetEnv("BACKEND_PASSWORD"),

		BackendTimeout: getEnvAsDuration("BACKEND_TIMEOUT_MS", 5000),
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return val
}

func getEnvAsDuration(key string, defaultMs int) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		ms, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("Invalid value for %s: %v", key, err)
		}
		return time.Duration(ms) * time.Millisecond
	}
	return time.Duration(defaultMs) * time.Millisecond
}
