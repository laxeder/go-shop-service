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

	route.Put("/user/password/:document", u.UpdateUserPassword)
	route.Put("/user/document/:document", u.UpdateUserDocument)
	route.Put("/user/restore/:document", u.RestoreUser)
	route.Delete("/user/:document", u.DeleteUser)
	route.Put("/user/:document", u.UpdateUser)
	route.Get("/user/:document", u.ShowUser)
	route.Post("/user", u.CreateUser)
	route.Patch("/user", u.ListUsers)

	route.Put("/address/document/:document", a.UpdateAddressDocument)
	route.Put("/address/restore/:document", a.RestoreAddress)
	route.Delete("/address/:document", a.DeleteAddress)
	route.Put("/address/:document", a.UpdateAddress)
	route.Get("/address/:document", a.ShowAddress)
	route.Post("/address", a.CreateAddress)
	route.Patch("/address", a.ListAddress)

	route.Put("/account/document/:document", c.UpdateAccountDocument)
	route.Put("/account/restore/:document", c.RestoreAccount)
	route.Delete("/account/:document", c.DeleteAccount)
	route.Put("/account/:document", c.UpdateAccount)
	route.Get("/account/:document", c.ShowAccount)
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
