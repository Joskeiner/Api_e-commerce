package v1

import (
	"errors"

	"github.com/Joskeiner/Api_e-commerce/internal/app/shop/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

// ShopControllerV1 is a struct for  version 1 of ShopController
type ShopControllerV1 struct {
	uc     domain.ShopUseCase
	server *http.Http
}

// new creates a new instance of ShopControllerV1
func New(uc domain.ShopUseCase, server *http.Http) *ShopControllerV1 {
	return &ShopControllerV1{
		uc,
		server,
	}
}

// GetAll handles GET /shops request
func (sc *ShopControllerV1) GetAll(ctx *fiber.Ctx) error {
	var q shopQuery

	if err := ctx.QueryParser(&q); err != nil {
		sc.server.Logger.Error("faild to parse query", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)

	}
	if err := sc.server.Validate.Struct(&q); err != nil {
		sc.server.Logger.Error("faild to Validate query ", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}
	res, err := sc.uc.GetAll(ctx.Context(), q.Page, q.Limit, q.Name)
	if err != nil {
		sc.server.Logger.Error("faild to get shops", "error", err)
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)

	}

	var shops []shopResponse
	for _, shop := range res {
		shops = append(shops, *NewShopRespose(shop))
	}

	rsp := &shopResponsePaginated{
		Shops: shops,
		Page:  q.Page,
		Limit: q.Limit,
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, err, rsp)
}

// GetUserShop handles Get /shops/myshop request
func (sc *ShopControllerV1) GetUserShop(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	res, err := sc.uc.GetUserShop(ctx.Context(), userID)
	if err != nil {
		sc.server.Logger.Error("faild to get shop", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}
	rsp := NewShopRespose(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, err, rsp)
}

// GetByID handles Get /shops/:id request
func (sc *ShopControllerV1) GetByID(ctx *fiber.Ctx) error {
	var p shopParam

	if err := ctx.ParamsParser(&p); err != nil {
		sc.server.Logger.Error("faild to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}
	if err := sc.server.Validate.Struct(&p); err != nil {
		sc.server.Logger.Error("faild to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}
	res, err := sc.uc.GetByID(ctx.Context(), p.ID)
	if err != nil {
		sc.server.Logger.Error("faild to get shop", "error", err)
		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}
	rsp := NewShopRespose(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// Update handels PUT /shops/:id resquest
func (sc *ShopControllerV1) Update(ctx *fiber.Ctx) error {
	var shopUR updateShopResquest

	if err := ctx.BodyParser(&shopUR); err != nil {
		sc.server.Logger.Error("faild to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}

	var p shopParam
	if err := ctx.ParamsParser(&p); err != nil {
		sc.server.Logger.Error("faild to parse path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}
	if err := sc.server.Validate.Struct(&p); err != nil {
		sc.server.Logger.Error("faild to validate path parameter", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDGETDATA, err, nil)
	}
	userID := ctx.Locals("user_id").(uint)

	if err := sc.uc.Update(ctx.Context(), userID, p.ID); err != nil {
		sc.server.Logger.Error("faild to update shop", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}
		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}
	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDPUTDATA, nil, nil)
}
