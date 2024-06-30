package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const ()

// Config ...
type Config struct {
	Environment   string // develop, staging, production
	UserServiceHost string
	UserServicePort string
	ManaServiceHost string
	ManaServicePort string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "prod"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))


	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("User_GRPC_PORT", "8082"))


	c.ManaServiceHost = cast.ToString(getOrReturnDefault("Mana_SERVICE_HOST", "localhost"))
	c.ManaServicePort = cast.ToString(getOrReturnDefault("Mana_GRPC_PORT", "8083"))



	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}

	return os.Getenv(key)
}
