package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// error de rota nõ encontrada
func NotFound(ctx *fiber.Ctx) error {
	return response.Ctx(ctx).Result(response.Error(404, "BLC000", "Recurso não encontrado."))
}
