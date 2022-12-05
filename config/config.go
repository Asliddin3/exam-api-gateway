package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment             string
	CustomerServiceHost     string
	CustomerServicePort     int
	ReviewServiceHost       string
	ReviewServicePort       int
	PostServiceHost         string
	PostServicePort         int
	CtxTimeout              int
	LogLevel                string
	HTTPPort                string
	SignKey                 string
	PostgresHost            string
	PostgresPort            int
	PostgresUser            string
	PostgresDB              string
	PostgresPassword        string
	AuthConfigPath          string
	SigninKey               string
	KAFKA_BROKER_ID         string
	KAFKA_ZOOKEEPER_CONNECT string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8070"))
	c.CustomerServiceHost = cast.ToString(getOrReturnDefault("CUSTOMER_SERVICE_HOST", "customer-servis"))
	c.CustomerServicePort = cast.ToInt(getOrReturnDefault("CUSTOMER_SERVICE_PORT", 8810))
	c.ReviewServiceHost = cast.ToString(getOrReturnDefault("REVIEW_SERVICE_HOST", "review-servise"))
	c.ReviewServicePort = cast.ToInt(getOrReturnDefault("REVIEW_SERVICE_PORT", 8840))
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post-service"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 8820))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "asliddin2001"))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "api"))
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_PATH", "./config/auth.conf"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 13))
	c.SignKey = cast.ToString(getOrReturnDefault("SECRET_KEY", "supersecret"))
	c.SigninKey = cast.ToString(getOrReturnDefault("SIGNIN_KEY", "supersecret"))

	return c

}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
