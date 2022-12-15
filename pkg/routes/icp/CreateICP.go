package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

// cria um CDU para uma usuário na base de dados
func CreateICP(ctx *fiber.Ctx) error {
	// var log = logger.New()

	// body := ctx.Body()
	// document := ctx.Params("document")

	// // converte json para struxt
	// icpBody, err := icp.New(body)
	// if err != nil {
	// 	log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS238", "O formado dos dados envidados está incorreto."))
	// }

	// spew.Dump(document, icpBody)
	// // valia se os documentos combinam
	// if str.DocumentPad(document) != str.DocumentPad(icpBody.Document) {
	// 	log.Error().Msg("O documento passado não corresponde ao documento do usuário." + icpBody.Document)
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS239", "O documento passado não corresponde ao documento do usuário. "+icpBody.Document))
	// }

	// // valida os campos de entrada
	// checkICP := icpBody.Valid()
	// if checkICP.Status != 200 {
	// 	return response.Ctx(ctx).Result(checkICP)
	// }

	// //!##################################################################################################################//
	// //! VERIFICA SE O ICP PARA ESTE USUÁRIO EXISTE NA BASE DE DADOS 													 //
	// //!##################################################################################################################//
	// icpData, err := icp.Repository().GetByDocument(icpBody.Document)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Erro ao tentar encontrar um icp para o usuário " + document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS240"))
	// }

	// // verifica se um documento e um serial number de certificado existe para este usuário
	// if len(icpData.Document) > 0 && len(icpData.SerialNumber) > 0 {
	// 	log.Error().Msg("Este certificado já existe na nossa base de dados.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS241", "Este certificado já existe na nossa base de dados."))
	// }

	// // ! #############################################################################################################################
	// // ! CRIA UMA ICP E ARMAZENA NA BASE DE DADOS
	// // ! #############################################################################################################################
	// icpBody.NewUuid()

	// icpBody.Validate = date.BRToUTC(icpBody.Validate)
	// icpBody.Status = icp.Enabled
	// icpBody.CreatedAt = date.NowUTC()
	// icpBody.UpdatedAt = date.NowUTC()

	// // armazena a ICP na base de dados de CDUS
	// err = icp.Repository().Save(icpBody)
	// if err != nil {
	// 	log.Error().Err(err).Msg("error ao tentar salvar a ICP do usuário." + icpBody.Document)
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS242"))
	// }

	return response.Ctx(ctx).Result(response.Success(201))
}
