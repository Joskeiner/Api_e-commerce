package app

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/config"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/log"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
)

// Run is the entrypoint of the application , dependecies are injected here
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

	server := http.New(cfg.Http, log)

	// dependency injection
	log.Info("strating the aplication", "name", cfg.App.Name, "enviroment", cfg.App.Env)

	if err := server.Start(); err != nil {
		log.Fatal("faild to start the server", "error", err)
	}
}
