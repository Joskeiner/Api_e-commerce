package http

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/config"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Http struct {
	App      *fiber.App
	Validate *validator.Validate
	Logger   log.Logger
	cfg      *config.Http
}

func New(cfg *config.Http, logger log.Logger) *Http {
	app := fiber.New()

	fiberLogger := customLogger(logger)

	cors := cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
	})

	app.Use(fiberLogger, cors)

	validate := validator.New()

	return &Http{
		app,
		validate,
		logger,
		cfg,
	}
}

func (h *Http) Start() error {
	listenAddr := h.cfg.Url + ":" + h.cfg.Port
	return h.App.Listen(listenAddr)
}

func customLogger(l log.Logger) fiber.Handler {
	format := "${status} | ${method} | ${path} | ${protocol} | ${ip} | ${latency} | ${ua}\n"

	return logger.New(logger.Config{
		Format:        format,
		DisableColors: true,
		Done: func(c *fiber.Ctx, logString []byte) {
			l.Info("request", "log", string(logString))
		},
	})
}
