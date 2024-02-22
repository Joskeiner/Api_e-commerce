package address

import (
	v1 "github.com/Joskeiner/Api_e-commerce/internal/app/address/controller/http/v1"
	"github.com/Joskeiner/Api_e-commerce/internal/app/address/repository"
	"github.com/Joskeiner/Api_e-commerce/internal/app/address/usecase"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/database"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
)

func New(db database.DB, server *http.Http, token token.Token) {
	addRepo := repository.New(db)
	addrUseCase := usecase.New(addRepo)
	addrV1 := v1.New(addrUseCase, server)

	routeV1 := server.App.Group("v1/users/address")
	addrV1.InitRoutes(routeV1, token)
}
