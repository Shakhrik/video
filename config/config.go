package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	// PostgresHost         string
	// PostgresPort         int
	// PostgresUser         string
	// PostgresPassword     string
	// PostgresDatabase     string
	// PostgresTestDatabase string

	LogLevel string
	HttpPort string
	AuthURL  string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found", err)
	}
	config := Config{}
	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8002"))
	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
