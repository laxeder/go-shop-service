package routes

import (
	"github.com/gofiber/fiber/v2"
	c "github.com/laxeder/go-shop-service/pkg/routes/account"
	a "github.com/laxeder/go-shop-service/pkg/routes/address"
	p "github.com/laxeder/go-shop-service/pkg/routes/product"
	u "github.com/laxeder/go-shop-service/pkg/routes/user"
)

func ApiV1(app *fiber.App) {

	route := app.Group("/api/v1")

	route.Get("/health", Health)
	route.Get("/redis/health", RedisHealth)

	route.Put("/user/password/:uuid", u.UpdateUserPassword)
	route.Put("/user/document/:uuid", u.UpdateUserDocument)
	route.Put("/user/restore/:uuid", u.RestoreUser)
	route.Delete("/user/:uuid", u.DeleteUser)
	route.Put("/user/:uuid", u.UpdateUser)
	route.Get("/user/:uuid", u.ShowUser)
	route.Post("/user", u.CreateUser)
	route.Patch("/user", u.ListUsers)

	route.Put("/address/restore/:uid", a.RestoreAddress)
	route.Delete("/address/:uid", a.DeleteAddress)
	route.Put("/address/:uid", a.UpdateAddress)
	route.Get("/address/:uid", a.ShowAddress)
	route.Post("/address", a.CreateAddress)
	route.Patch("/address", a.ListAddress)

	route.Put("/account/restore/:uid", c.RestoreAccount)
	route.Delete("/account/:uid", c.DeleteAccount)
	route.Put("/account/:uid", c.UpdateAccount)
	route.Get("/account/:uid", c.ShowAccount)
	route.Post("/account", c.CreateAccount)
	route.Patch("/account", c.ListAccounts)

	route.Put("/product/restore/:uid", p.RestoreProduct)
	route.Put("/product/uid/:uid", p.UpdateProductUid)
	route.Delete("/product/:uid", p.DeleteProduct)
	route.Put("/product/:uid", p.UpdateProduct)
	route.Get("/product/:uid", p.ShowProduct)
	route.Post("/product", p.CreateProduct)
	route.Patch("/product", p.ListProducts)

}

func ErrorNotFound(app *fiber.App) {
	app.Use("/", NotFound)
}
