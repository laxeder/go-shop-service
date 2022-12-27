package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	j "github.com/laxeder/go-shop-service/pkg/modules/jwt"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func Claims(ctx *fiber.Ctx) *j.Claims {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	cl := &j.Claims{}

	return cl.InjectMap(claims)
}

func UserClaims(ctx *fiber.Ctx) *user.User {

	claims := Claims(ctx)

	u := &user.User{}

	return u.InjectMap(claims.Data)
}
