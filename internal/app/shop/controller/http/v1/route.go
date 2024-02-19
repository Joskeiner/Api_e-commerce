package v1

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/middleware"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func (sc *ShopControllerV1) InitRoutes(routerV1 fiber.Router, token token.Token) {
	routerV1.Get("/", sc.GetAll)
	routerV1.Get("/myshop", middleware.AuthMiddleware(token), sc.GetUserShop)
	routerV1.Get("/:id", sc.GetByID)
	routerV1.Get("/:id", middleware.AuthMiddleware(token), sc.Update)
}
