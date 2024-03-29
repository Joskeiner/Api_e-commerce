package v1

import (
	"errors"
	"time"

	"github.com/Joskeiner/Api_e-commerce/internal/app/user/domain"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/server/http"
	"github.com/gofiber/fiber/v2"
)

type UserControllerV1 struct {
	uc     domain.UserUsecase
	server *http.Http
}

// NewUserControllerV1 creates a new instance of UserControllerV1
func New(uc domain.UserUsecase, server *http.Http) *UserControllerV1 {
	return &UserControllerV1{
		uc,
		server,
	}
}

// GetByID handles GET /users request
func (uc *UserControllerV1) GetByID(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)
	res, err := uc.uc.GetByID(ctx.Context(), userID)
	if err != nil {
		uc.server.Logger.Error("failed to get user", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDGETDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDGETDATA, err, nil)
	}

	rsp := NewUserResponse(*res)

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDGETDATA, nil, rsp)
}

// Update handles PUT /users request
func (uc *UserControllerV1) Update(ctx *fiber.Ctx) error {
	var req updateUserRequest

	if err := ctx.BodyParser(&req); err != nil {
		uc.server.Logger.Error("failed to parse request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	if err := uc.server.Validate.Struct(&req); err != nil {
		uc.server.Logger.Error("failed to validate request body", "error", err)
		return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
	}

	userID := ctx.Locals("user_id").(uint)

	var (
		birthDate time.Time
		err       error
	)

	if req.BirthDate != "" {
		birthDate, err = helper.ParseTime(req.BirthDate)
		if err != nil {
			uc.server.Logger.Error("failed to parse request body", "error", err)
			return helper.Response(ctx, fiber.StatusBadRequest, false, helper.FAILEDPUTDATA, err, nil)
		}
	}

	user := &domain.User{
		ID:          userID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		BirthDate:   birthDate,
		About:       req.About,
		Job:         req.Job,
		ProvinceID:  req.ProvinceID,
		CityID:      req.CityID,
	}

	if err := uc.uc.Update(ctx.Context(), user); err != nil {
		uc.server.Logger.Error("failed to update user", "error", err)

		if errors.Is(err, helper.ErrDataNotFound) {
			return helper.Response(ctx, fiber.StatusNotFound, false, helper.FAILEDPUTDATA, err, nil)
		}

		return helper.Response(ctx, fiber.StatusInternalServerError, false, helper.FAILEDPUTDATA, err, nil)
	}

	return helper.Response(ctx, fiber.StatusOK, true, helper.SUCCEEDPUTDATA, nil, nil)
}
