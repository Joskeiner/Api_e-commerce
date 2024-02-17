package v1

import "github.com/gofiber/fiber/v2"

// InitRoutes register all routes for version 1
func (cc *CityControllerV1) InitRoutes(routeV1 fiber.Router) {
	routeV1.Get("/:prov_id", cc.GetAll)
	routeV1.Get("/:prov_id/cities/:city_id", cc.GetByID)
}
