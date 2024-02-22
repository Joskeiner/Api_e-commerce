package v1

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/middleware"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func (ac *AddressControllerV1) InitRoutes(routeV1 fiber.Router, token token.Token) {
	routeV1.Use(middleware.AuthMiddleware(token))

	routeV1.Post("/", ac.Create)
	routeV1.Get("/", ac.GetAll)
	routeV1.Get("/:id", ac.GetByID)
	routeV1.Put("/:id", ac.Update)
	routeV1.Delete("/:id", ac.Delete)
}
