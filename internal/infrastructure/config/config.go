// Package config loads application configuration from the environment into typed structs.
package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
	LLM      LLMConfig
	Engine   EngineConfig
	Admin    AdminConfig

	// RequireEventToken, when true (and an admin key is set), requires an app token on event requests.
	RequireEventToken bool
}

type AppConfig struct {
	ServiceName string
	Environment string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	MongoURL      string
	MongoDatabase string
}

type CacheConfig struct {
	RedisURL string
}

type LLMConfig struct {
	GeminiKey string
}

type EngineConfig struct {
	MaxIterations int
}

// AdminConfig secures the management API. When APIKey is empty the management routes are not mounted.
type AdminConfig struct {
	APIKey string
}

// Load reads configuration from the environment, applying defaults for everything except secrets.
func Load() *Config {
	return &Config{
		App: AppConfig{
			ServiceName: getenv("SERVICE_NAME", "eywa-starter"),
			Environment: getenv("ENVIRONMENT", "lcl"),
		},
		Server: ServerConfig{Port: getenv("PORT", "8080")},
		Database: DatabaseConfig{
			MongoURL:      getenv("MONGO_URL", "mongodb://localhost:27017"),
			MongoDatabase: getenv("MONGO_DATABASE", "eywa_starter"),
		},
		Cache: CacheConfig{RedisURL: getenv("REDIS_URL", "redis://localhost:6379")},
		LLM:   LLMConfig{GeminiKey: mustEnv("GEMINI_API_KEY")},
		Engine: EngineConfig{
			MaxIterations: getenvInt("MAX_ITERATIONS", 5),
		},
		Admin:             AdminConfig{APIKey: os.Getenv("ADMIN_API_KEY")},
		RequireEventToken: os.Getenv("REQUIRE_EVENT_TOKEN") == "true",
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getenvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s is required (copy .env.example to .env and fill it in)", key)
	}
	return v
}
