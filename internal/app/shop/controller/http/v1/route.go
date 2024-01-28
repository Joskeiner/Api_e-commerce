package v1

import (
	"github.com/gofiber/fiber/v2"
)

func (sc *ShopControllerV1) InitRoutes(routerV1 fiber.Router) {
	routerV1.Get("/", sc.GetAll)
	// falta crear un middleware que tendra los datos que necesita la url para funcionar  sin problemas
	routerV1.Get("/myshop", sc.GetUserShop)
	routerV1.Get("/:id", sc.GetByID)
	// falta crear un middleware que tendra los datos que necesita la url para funcionar  sin problemas
	routerV1.Get("/:id", sc.Update)
}
