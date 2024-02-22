package v1

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/middleware"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
)

// InitRoutes registers all routes for version 1
func (uc *UserControllerV1) InitRoutes(routeV1 fiber.Router, token token.Token) {
	routeV1.Use(middleware.AuthMiddleware(token))

	routeV1.Get("/", uc.GetByID)
	routeV1.Put("/", uc.Update)
}
