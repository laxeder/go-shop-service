package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Claims(ctx *fiber.Ctx) jwt.MapClaims {
	return ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
}
