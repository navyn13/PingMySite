package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load() //loads key values from env and load them into system environment variables
	return Config{
		Port: getEnv("PORT", "4000"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
