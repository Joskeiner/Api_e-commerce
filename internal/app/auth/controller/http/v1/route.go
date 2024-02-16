package v1

import "github.com/gofiber/fiber/v2"

func (ac *AuthControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Post("register", ac.Register)
}
