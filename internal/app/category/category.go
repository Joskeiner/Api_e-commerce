package category

import (
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"

	v1 "github.com/Joskeiner/Api_e-commerce/internal/app/category/controller/http/v1"
	"github.com/Joskeiner/Api_e-commerce/internal/app/category/repository"
	"github.com/Joskeiner/Api_e-commerce/internal/app/category/usecase"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
)

// New injects the dependencies of category package
func New(db database.DB, server *http.Http, token token.Token) {
	categoryRepo := repository.New(db)
	categoryUsecase := usecase.New(categoryRepo)
	categoryV1 := v1.New(categoryUsecase, server)

	routeV1 := server.App.Group("/v1/categories")
	categoryV1.InitRoutes(routeV1, token)
}
