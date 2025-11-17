package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env string `env:"ENV"`

	App     AppConfig
	HTTP    HTTPConfig
	Logger  LoggerConfig
	JWT     JWTConfig
	PG      PostgresConfig
	Swagger SwaggerConfig
}

type AppConfig struct {
	Name    string `envconfig:"APP_NAME"`
	Version string `envconfig:"APP_VERSION"`
}

type HTTPConfig struct {
	Port int `envconfig:"HTTP_PORT"`
}

type LoggerConfig struct {
	Level string `envconfig:"LOG_LEVEL"`
}

type JWTConfig struct {
	AccessSecret  string `envconfig:"ACCESS_SECRET"`
	RefreshSecret string `envconfig:"REFRESH_SECRET"`
}

type PostgresConfig struct {
	PoolMax int    `envconfig:"PG_POOL_MAX"`
	URL     string `envconfig:"PG_URL"`
}

type SwaggerConfig struct {
	Enabled bool `envconfig:"SWAGGER_ENABLED"`
}

func Load() (*Config, error) {
	if err := godotenv.Load("./config.env"); err != nil {
		return nil, err
	}
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
