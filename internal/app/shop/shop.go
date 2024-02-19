package shop

import (
	v1 "github.com/Joskeiner/Api_e-commerce/internal/app/shop/controller/http/v1"
	"github.com/Joskeiner/Api_e-commerce/internal/app/shop/repository"
	"github.com/Joskeiner/Api_e-commerce/internal/app/shop/usecase"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
)

// new injects the dependecies of shop package
func New(db database.DB, server *http.Http, token token.Token) {
	shopRepo := repository.New(db)
	shopUsecase := usecase.New(shopRepo)
	shopV1 := v1.New(shopUsecase, server)

	routerV1 := server.App.Group("/v1/shops")
	shopV1.InitRoutes(routerV1, token)
}
