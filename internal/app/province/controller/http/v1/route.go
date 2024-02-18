package v1

import "github.com/gofiber/fiber/v2"

// InintRoutes registyer all riutes for version 1
func (pc *provinceControllerV1) InintRoutes(routeV1 fiber.Router) {
	routeV1.Get("/:id", pc.GetById)
	routeV1.Get("/", pc.GetAll)
}
