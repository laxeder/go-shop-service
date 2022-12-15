package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func Login(ctx *fiber.Ctx) error {

	// var log = logger.New()

	// body := ctx.Body()

	// loginBody, err := auth.LogIn(body)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Os campos enviados estão incorretos.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS002", "Os campos enviados estão incorretos."))
	// }

	// //!##################################################################################################################//
	// //! VERIFICA SE O DOCUMENTO DO USUÁRIO EXISTE NA BASE DE DADOS 														                           //
	// //!##################################################################################################################//
	// userData, result := shared.LoginCheck(loginBody)
	// if result.Status != 200 {
	// 	return response.Ctx(ctx).Result(result)
	// }

	// //!##################################################################################################################//
	// //! CRIA UMA NOVA SESSÃO PARA O USUÁRIO													 							 																		   //
	// //!##################################################################################################################//

	// // jwtUser := jwt.User{Uuid: userData.Uuid, Name: userData.FullName}
	// jwtData := jwt.New()
	// jwtData.Claims.Data = userData

	// return response.Ctx(ctx).Result(response.Success(200, jwtData.Result()))
	return response.Ctx(ctx).Result(response.Success(200))

}
