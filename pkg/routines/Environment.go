package routines

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func EnvironmentLoad(app *fiber.App) {

	// busca o ambiente
	env := os.Getenv("ENVIRONMENT")

	envFile := "./.env.prod"
	if env == "dev" {
		envFile = "./.env.dev"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("As variáveis de ambiente nao foram encontradas: %v", err)
		return
	}
}
