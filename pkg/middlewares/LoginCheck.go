package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func LoginCkeck(ctx *fiber.Ctx) error {

	// var log = logger.New()

	// body := ctx.Body()

	// loginBody, err := auth.LogIn(body)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Os campos enviados estão incorretos.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "BLC002", "Os campos enviados estão incorretos."))
	// }

	// checkLogin := loginBody.Valid()
	// if checkLogin.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkLogin)
	// }

	// _, result := shared.LoginCheck(loginBody)
	// if result.Status != 200 {
	// 	return response.Ctx(ctx).Result(result)
	// }

	return ctx.Next()
}
