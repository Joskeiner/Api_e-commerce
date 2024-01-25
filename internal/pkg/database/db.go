package database

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/config"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"gorm.io/gorm"
)

type DB interface {
	Migrate() error
	// DB returns the underlying gorm.DB connection
	DB() *gorm.DB

	Close() error
}

func New(cfg *config.Databese) (DB, error) {
	switch cfg.Conn {
	case "postgres":
		return newPostgres(cfg)
	}
	return nil, helper.ErrUnsupportedDriver
}
