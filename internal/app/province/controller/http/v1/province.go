package v1

import (
	"github.com/Joskeiner/Api_e-commerce/internal/app/province/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// provinceControllerV1 is a struct for version 1 of provinceControllerV1
type provinceControllerV1 struct {
	uc     domain.ProvinceUseCase
	server *http.Http
}

// New creates a new instance of provinceControllerV1
func New(uc domain.ProvinceUseCase, server *http.Http) *provinceControllerV1 {
	return &provinceControllerV1{
		uc,
		server,
	}
}

func (pc *provinceControllerV1) GetAll(ctx *fiber.Ctx) error {
	res, err := pc.uc.GetAll()
	if err != nil {
		pc.server.Logger.Error("faild to get provinces ", "error", err)
		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	var rsp []provinceResponse
	for _, province := range res {
		rsp = append(rsp, *NewProvinceResponse(province))
	}
	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// GetByID is function to get province by id
func (pc *provinceControllerV1) GetById(ctx *fiber.Ctx) error {
	var province provinceParam
	if err := ctx.ParamsParser(&province); err != nil {
		pc.server.Logger.Error("faild to get province ", "error", err)
	}
	res, err := pc.uc.GetByID(province.ID)
	if err != nil {
		pc.server.Logger.Error("faild to get province", "error", err)
		if err == helper.ErrDataNotFound {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewProvinceResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}
