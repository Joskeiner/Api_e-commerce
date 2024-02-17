package v1

import (
	"github.com/Joskeiner/Api_e-commerce/internal/app/city/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// CityControllerV1 is a struct for version 1 of CityController
type CityControllerV1 struct {
	cu     domain.CityUseCase
	server *http.Http
}

// New creates a new instance of CityControllerV1
func New(cu domain.CityUseCase, server *http.Http) *CityControllerV1 {
	return &CityControllerV1{
		cu,
		server,
	}
}

// GetAll is a function to get all cities
func (cc *CityControllerV1) GetAll(ctx *fiber.Ctx) error {
	var cityP citiesParam

	if err := ctx.ParamsParser(&cityP); err != nil {
		cc.server.Logger.Error("faild to parse path parameter ", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}
	res, err := cc.cu.GetAll(cityP.ProvinceID)
	if err != nil {
		cc.server.Logger.Error("faild to get cities ", "error", err)
		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}
	var rsp []cityResponse
	for _, city := range res {
		rsp = append(rsp, *NewCityResponse(city))
	}
	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetByID is a function to get by ProvinceID
func (cc *CityControllerV1) GetByID(ctx *fiber.Ctx) error {
	var cityP cityParam
	if err := ctx.ParamsParser(&cityP); err != nil {
		cc.server.Logger.Error("faild to parse path parameter ", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}
	res, err := cc.cu.GetByID(cityP.ProvinceID, cityP.ID)
	if err != nil {
		cc.server.Logger.Error("faild to get city", "error", err)

		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}
	rsp := NewCityResponse(*res)
	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}
