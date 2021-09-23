package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	LogLevel string

	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8000"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
