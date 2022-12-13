package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// executa cógido
func Hooks(ctx *fiber.Ctx) error {

	now := time.Now().UTC()
	result := fmt.Sprintf("Hoock: %s", now.Format("2006-01-02 15:04:05"))

	return response.Ctx(ctx).Result(response.Success(200, result))
}
