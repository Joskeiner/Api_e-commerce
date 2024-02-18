package province

import (
	v1 "github.com/Joskeiner/Api_e-commerce/internal/app/province/controller/http/v1"
	"github.com/Joskeiner/Api_e-commerce/internal/app/province/repository"
	"github.com/Joskeiner/Api_e-commerce/internal/app/province/usecase"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
)

func New(server *http.Http) {
	provRepo := repository.New()
	provUseCase := usecase.New(provRepo)
	provinceV1 := v1.New(provUseCase, server)

	routeV1 := server.App.Group("/v1/provinces")
	provinceV1.InintRoutes(routeV1)
}
