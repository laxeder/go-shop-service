package routes

import (
	"github.com/gofiber/fiber/v2"
	acc "github.com/laxeder/go-shop-service/pkg/routes/account"
	addr "github.com/laxeder/go-shop-service/pkg/routes/address"
	ctg "github.com/laxeder/go-shop-service/pkg/routes/category"
	prod "github.com/laxeder/go-shop-service/pkg/routes/product"
	usr "github.com/laxeder/go-shop-service/pkg/routes/user"
)

func ApiV1(app *fiber.App) {

	route := app.Group("/api/v1")

	//? ********** Rotas do servidor **********

	route.Get("/health", Health)
	route.Get("/redis/health", RedisHealth)

	//? *************** Rotas do usuário ***************

	route.Put("/user/password/:uuid", usr.UpdateUserPassword)
	route.Put("/user/document/:uuid", usr.UpdateUserDocument)
	route.Patch("/user/:uuid", usr.RestoreUser)
	route.Delete("/user/:uuid", usr.DeleteUser)
	route.Put("/user/:uuid", usr.UpdateUser)
	route.Get("/user/:uuid", usr.ShowUser)
	route.Post("/user", usr.CreateUser)
	route.Get("/users", usr.ListUsers)

	//? *************** Rotas do endereço ***************

	route.Patch("/address/:uid", addr.RestoreAddress)
	route.Delete("/address/:uid", addr.DeleteAddress)
	route.Put("/address/:uid", addr.UpdateAddress)
	route.Get("/address/:uid", addr.ShowAddress)
	route.Post("/address", addr.CreateAddress)
	route.Get("/adresses", addr.ListAddress)

	//? *************** Rotas da conta ***************

	route.Patch("/account/:uid", acc.RestoreAccount)
	route.Delete("/account/:uid", acc.DeleteAccount)
	route.Put("/account/:uid", acc.UpdateAccount)
	route.Get("/account/:uid", acc.ShowAccount)
	route.Post("/account", acc.CreateAccount)
	route.Get("/accounts", acc.ListAccounts)

	//? *************** Rotas do produto ***************

	route.Patch("/product/:uid", prod.RestoreProduct)
	route.Delete("/product/:uid", prod.DeleteProduct)
	route.Put("/product/:uid", prod.UpdateProduct)
	route.Get("/product/:uid", prod.ShowProduct)
	route.Post("/product", prod.CreateProduct)
	route.Get("/products", prod.ListProducts)

	//? *************** Rotas da categoria ***************

	route.Patch("/category/:code", ctg.RestoreCategory)
	route.Delete("/category/:code", ctg.DeleteCategory)
	route.Put("/category/:code", ctg.UpdateCategory)
	route.Get("/category/:code", ctg.ShowCategory)
	route.Post("/category", ctg.CreateCategory)
	route.Get("/categories", ctg.ListCategories)

}

func ErrorNotFound(app *fiber.App) {
	app.Use("/", NotFound)
}
