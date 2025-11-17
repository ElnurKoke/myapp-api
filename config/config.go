package config

import "github.com/kelseyhightower/envconfig"

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
	Name    string `env:"APP_NAME"`
	Version string `env:"APP_VERSION"`
}

type HTTPConfig struct {
	Port int `env:"HTTP_PORT"`
}

type LoggerConfig struct {
	Level string `env:"LOG_LEVEL"`
}

type JWTConfig struct {
	AccessSecret  string `env:"ACCESS_SECRET"`
	RefreshSecret string `env:"REFRESH_SECRET"`
}

type PostgresConfig struct {
	PoolMax int    `env:"PG_POOL_MAX"`
	URL     string `env:"PG_URL"`
}

type SwaggerConfig struct {
	Enabled bool `env:"SWAGGER_ENABLED"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
