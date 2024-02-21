package middleware

import (
	"strings"

	"github.com/Joskeiner/Api_e-commerce/internal/pkg/helper"
	"github.com/Joskeiner/Api_e-commerce/internal/pkg/token"
	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationType      = "bearer"
)

func AuthMiddleware(token token.Token) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.Get(authorizationHeaderKey)

		isEmpty := len(authorizationHeader) == 0
		if isEmpty {
			err := helper.ErrEmptyAuthorizationHeader
			return helper.Response(ctx, 401, false, helper.UNAUTHORIZED, err, nil)
		}
		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := helper.ErrInvalidAuthorizationHeader
			return helper.Response(ctx, 401, false, helper.UNAUTHORIZED, err, nil)
		}
		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != authorizationType {
			err := helper.ErrInvalidAuthorizationType
			return helper.Response(ctx, 401, false, helper.UNAUTHORIZED, err, nil)
		}
		accesToken := fields[1]
		payload, err := token.Verify(accesToken)
		if err != nil {
			return helper.Response(ctx, 401, false, helper.UNAUTHORIZED, err, nil)
		}
		ctx.Locals("user_id", payload.UserID)
		ctx.Locals("is_admin", payload.IsAdmin)

		return ctx.Next()
	}
}

// adminMiddleware is a middleware to check if the user is an admin
func AdminMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		isAdmin := ctx.Locals("is_admin").(bool)

		if !isAdmin {
			err := helper.ErrInsufficientPermission
			return helper.Response(ctx, 403, false, helper.FORBIDDEN, err, nil)
		}
		return ctx.Next()
	}
}
