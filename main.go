package main

import (
	_ "github.com/laxeder/go-shop-service/docs" // load API Docs files (Swagger)
	"github.com/laxeder/go-shop-service/pkg"
	"github.com/laxeder/go-shop-service/pkg/middlewares"
	"github.com/laxeder/go-shop-service/pkg/routes"
	"github.com/laxeder/go-shop-service/pkg/routines"
)

func main() {

	server := &pkg.Server{}
	server.New()
	server.Routines(routines.EnvironmentLoad)
	server.Routines(routines.Run)
	server.Middlewares(middlewares.Logger)
	server.Middlewares(middlewares.Cors)
	server.Routes(routes.ApiV1)
	server.Routes(routes.Swagger)
	server.Routes(routes.ErrorNotFound)

	defer server.Start()
}

// CRUD (create, read, update, delete)

/*
 ! Objetivo: criar plataforma E-commerc loja virtual

 ? Sistema de contatos: Contacts CRUD
 ? Sistema de login: Auth JWT

 ? Sistema de compras: Buy CRUD
 ? Sistema de favoritos (factory buy)
 ? Sistema de frete: Freight
 ? Sistema de estoque (factory product)

 ? Dashboard
 	* vendas do dia: produtos
	* balanço do dia (quanto deve - quanto tem)
	* lead do dia: lista
	* produtos despachados
	* produtos entregues
	* produtos em transito
	* produtos a despachar

 ? Sistema de pagamento: integração do sistem de pagamento

	* listar produto por categoria
	* rota --> listar produto por categoria
	* listar produto por ordem alfabetica
	* listar produto por ordem de preço
	* listar produto por data de criação

 ? PICTURE

	* name
	* mimetype
	* base64

  TODO retornar result com erro se não tiver permissão
  TODO routines atualizar permissões
 	pegar lista de permissão da account permissão passada ele contem

*/
