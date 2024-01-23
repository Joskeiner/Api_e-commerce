package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type JSONResponse struct {
	Status  bool     `json:"status"`
	Massage string   `json:"message"`
	Errors  []string `json:"errors"`
	Data    any      `json:"data"`
}

func Response(ctx *fiber.Ctx, code int, status bool, message string, err error, data any) error {
	var errMsgs []string
	if err != nil {
		errMsgs = parseError(err)
	}

	res := &JSONResponse{
		Status:  status,
		Massage: message,
		Errors:  errMsgs,
		Data:    data,
	}

	return ctx.Status(code).JSON(res)
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}
