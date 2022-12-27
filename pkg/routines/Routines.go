package routines

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/gmud"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

var log = logger.New()

func Run(app *fiber.App) {
	DocSwagger()
	defer GMUDPermissions()
	defer AnalyzeCode()
}

func GMUDPermissions() {
	gmud.AddPermissions()
}
