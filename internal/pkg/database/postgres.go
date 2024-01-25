package database

import (
	"fmt"
	"time"

	"github.com/Joskeiner/Api_e-commerce/internal/pkg/config"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	Conn *gorm.DB
	cfg  *config.Databese
}

// url connection
func newPostgres(cfg *config.Databese) (DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Name,
		cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	maxLifeTime := time.Duration(cfg.MaxLifeTime) * time.Second

	sqlDB.SetConnMaxLifetime(maxLifeTime)

	return &Postgres{
		db,
		cfg,
	}, nil
}

func (p *Postgres) Migrate() error {
	return p.Conn.AutoMigrate(
		dao.User{},
		dao.Address{},
		dao.Shop{},
		dao.Category{},
		dao.Product{},
		dao.ProductPhoto{},
		dao.ProductLog{},
		dao.Transaction{},
		dao.TransactionDetail{},
	)
}

// DB returns the underlying database connection.
func (p *Postgres) DB() *gorm.DB {
	return p.Conn
}

// close closes the database connection
func (p *Postgres) Close() error {
	sqlDB, err := p.Conn.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
