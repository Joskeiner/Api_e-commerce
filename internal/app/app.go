package app

import (
	"github.com/Joskeiner/Api_e-commerce/internal/app/address"
	"github.com/Joskeiner/Api_e-commerce/internal/app/auth"
	"github.com/Joskeiner/Api_e-commerce/internal/app/category"
	"github.com/Joskeiner/Api_e-commerce/internal/app/city"
	"github.com/Joskeiner/Api_e-commerce/internal/app/province"
	"github.com/Joskeiner/Api_e-commerce/internal/app/shop"
	"github.com/Joskeiner/Api_e-commerce/internal/app/user"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/config"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/log"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
)

func Run() {
	log, err := log.New()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := log.Close(); err != nil {
			log.Panic("faild to close the logger", "error", err)
		}
	}()

	cfgLoader, err := config.New()
	if err != nil {
		log.Fatal("faild to initialize the config provider", "error", err)
	}

	cfg, err := cfgLoader.Load()
	if err != nil {
		log.Fatal("faild to load the config", "error", err)
	}
	log.Info("succeed to load the config")

	db, err := database.New(cfg.Databese)
	if err != nil {
		log.Fatal("faild to initialize the database")
	}
	defer db.Close()
	if err := db.Migrate(); err != nil {
		log.Fatal("faild to migrate the database", "error", err)
	}
	log.Info("succed to migrate the database")
	// config token
	token, err := token.New(cfg.Token)
	if err != nil {
		log.Fatal("faild to initialize the token provider", "error", err)
	}

	server := http.New(cfg.Http, log)

	// dependency injection
	user.New(db, server, token)
	category.New(db, server, token)
	address.New(db, server, token)
	province.New(server)
	city.New(server)
	auth.New(db, server, token)
	shop.New(db, server, token)
	log.Info("strating the aplication", "name", cfg.App.Name, "enviroment", cfg.App.Env)

	if err := server.Start(); err != nil {
		log.Fatal("faild to start the server", "error", err)
	}
}
