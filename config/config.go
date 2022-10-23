package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment         string
	CustomerServiceHost string
	CustomerServicePort int
	ReviewServiceHost   string
	ReviewServicePort   int
	PostServiceHost     string
	PostServicePort     int
	CtxTimeout          int
	LogLevel            string
	HTTPPort            string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8070"))
	c.CustomerServiceHost = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_HOST", "localhost"))
	c.CustomerServicePort = cast.ToInt(getOrReturnDefault("PRODUCT_SERVICE_PORT", 8810))
	c.ReviewServiceHost = cast.ToString(getOrReturnDefault("STORE_SERVICE_HOST", "localhost"))
	c.ReviewServicePort = cast.ToInt(getOrReturnDefault("STORE_SERVICE_PORT", 8840))
	c.PostServiceHost = cast.ToString(getOrReturnDefault("User_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("User_SERVICE_PORT", 8820))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c

}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
