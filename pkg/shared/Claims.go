package shared

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
	"github.com/laxeder/go-shop-service/pkg/utils"
	j "github.com/laxeder/go-shop-service/pkg/utils/jwt"
)

func Claims(ctx *fiber.Ctx) *j.Claims {
	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	cl := &j.Claims{}

	return cl.InjectMap(claims)
}

func UserClaims(ctx *fiber.Ctx) (u *user.User, err error) {

	u = &user.User{}

	claims := Claims(ctx)

	err = utils.InjectMap(claims.Data.(fiber.Map), u)

	return
}
