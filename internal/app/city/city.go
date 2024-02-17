package city

import (
	v1 "github.com/Joskeiner/Api_e-commerce/internal/app/city/controller/http/v1"
	"github.com/Joskeiner/Api_e-commerce/internal/app/city/repository"
	"github.com/Joskeiner/Api_e-commerce/internal/app/city/usecase"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
)

// New injects the depencies of city package
func New(server *http.Http) {
	cityRepo := repository.New()
	cityUseCase := usecase.New(cityRepo)
	cityV1 := v1.New(cityUseCase, server)

	routerV1 := server.App.Group("/v1/city")
	cityV1.InitRoutes(routerV1)
}
