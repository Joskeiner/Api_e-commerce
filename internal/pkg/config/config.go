package config

import "time"

type Config interface {
	Load() (*Container, error)
}

type (
	Container struct {
		App      *App
		Http     *Http
		Databese *Databese
		// Token    *Token
	}
	App struct {
		Name string `mapstructure:"APP_NAME"`
		Env  string `mapstructure:"APP_ENV"`
	}

	Http struct {
		Url            string `mapstructure:"HTTP_URL"`
		Port           string `mapstructure:"HTTP_PORT"`
		AllowedOrigins string `mapstructure:"HTTP_ALLOWED_ORIGINS"`
	}
	Databese struct {
		Conn        string `mapstructure:"DB_CONNECTION"`
		Host        string `mapstructure:"DB_HOST"`
		Port        string `mapstructure:"DB_PORT"`
		Name        string `mapstructure:"POSTGRES_DB"`
		Username    string `mapstructure:"POSTGRES_USER"`
		Password    string `mapstructure:"POSTGRES_PASSWORD"`
		MaxLifeTime int    `mapstructure:"DB_MAX_LIFE_TIME"`
		MaxOpenConn int    `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
		MaxIdleConn int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	}
	Token struct {
		Type         string        `mapstructure:"TOKEN_TYPE"`
		SymmetricKey string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
		Duration     time.Duration `mapstructure:"TOKEN_DURATION"`
	}
)

func New() (Config, error) {
	return newViper()
}
