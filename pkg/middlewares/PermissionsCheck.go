package middlewares

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/shared"
)

func PermissionsCheck(permission string) func(ctx *fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {
		user := shared.Claims(ctx)

		spew.Dump(permission, user)

		return ctx.Next()
	}
}
