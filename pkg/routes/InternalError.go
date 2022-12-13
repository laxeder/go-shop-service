package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// rota do erro 500
func internalError(ctx *fiber.Ctx) error {
	return response.Ctx(ctx).Result(response.ErrorDefault("BLC001"))
}
