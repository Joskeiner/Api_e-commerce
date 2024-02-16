package auth

import (
	v1 "github.com/Joskeiner/Api_e-commerce/internal/app/auth/controller/http/v1"
	"github.com/Joskeiner/Api_e-commerce/internal/app/auth/repository"
	usercase "github.com/Joskeiner/Api_e-commerce/internal/app/auth/usecase"
	cityRepository "github.com/Joskeiner/Api_e-commerce/internal/app/city/repository"
	provinceRepository "github.com/Joskeiner/Api_e-commerce/internal/app/province/repository"
	userRepository "github.com/Joskeiner/Api_e-commerce/internal/app/user/repository"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
)

// New injects the dependencies of auth package
func New(db database.DB, server *http.Http, token token.Token) {
	authRepo := repository.New(db)
	userRepo := userRepository.New(db)
	cityRepo := cityRepository.New()
	provRepo := provinceRepository.New()
	authUsecase := usercase.New(authRepo, userRepo, cityRepo, provRepo, token)
	authV1 := v1.New(authUsecase, server)

	routeV1 := server.App.Group("/v1/auth")
	authV1.InitRoutes(routeV1)
}
