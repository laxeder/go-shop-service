package jwt

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Config() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	})
}
