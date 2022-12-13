package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Token(ctx *fiber.Ctx) *jwt.Token {
	return ctx.Locals("user").(*jwt.Token)
}
