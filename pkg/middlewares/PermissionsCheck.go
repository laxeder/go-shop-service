package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func PermissionsCheck(permission string) func(ctx *fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {

		// user := shared.UserClaims(ctx)

		// user.Permissions = append(user.Permissions, permission)

		// spew.Dump(user.Permissions, permission)

		// if !utils.Contains(user.Permissions, permission) {
		// 	return response.Ctx(ctx).Result(response.Error(400, "GSS002", "Vocẽ não tem permissão para acessar essa rota."))
		// }

		//TODO: Criar repositorio para pegar as permissão do user

		return ctx.Next()
	}
}
