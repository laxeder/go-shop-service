package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func ICPCheck(ctx *fiber.Ctx) error {
	// var log = logger.New()

	// document := ctx.Params("document")

	// icpData, err := icp.Repository().GetByDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Erro ao tentar encontra o certificado ICP.")
	// 	return response.Ctx(ctx).Result(response.ErrorDefault("GSS251"))
	// }

	// if len(icpData.Document) == 0 && len(icpData.SerialNumber) == 0 {
	// 	log.Error().Msg("Este uauário não possui certificado ICP para registo digital.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS252", "Este uauário não possui certificado ICP para registo digital."))
	// }

	// if icpData.Status == icp.Disabled {
	// 	log.Error().Msg("O certificado ICP deste usuário está Desbilitado por tempo indeterminado.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS253", "O certificado ICP deste usuário está Desbilitado pro tempo indeterminado."))
	// }

	// if icpData.Status == icp.Expired {
	// 	log.Error().Msg("O certificado ICP deste  usuário está expirado.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS254", "O certificado ICP deste usuário está expirado."))
	// }

	// // caso o certificado esteja expirado e o status seja diferente de expirado, atualiza o database
	// if time.Now().Unix() > date.UTCToTime(icpData.Validate).Unix() {

	// 	icpData, err := icp.Repository().GetByDocument(document)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("Erro ao tentar encontrar um icp para o  usuário " + document)
	// 		return response.Ctx(ctx).Result(response.ErrorDefault("GSS255"))
	// 	}

	// 	icpData.Status = icp.Expired
	// 	icpData.UpdatedAt = date.NowUTC()

	// 	err = icp.Repository().Update(icpData)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("error ao tentar atualizar o ICP do usuário." + document)
	// 		return response.Ctx(ctx).Result(response.ErrorDefault("GSS256"))
	// 	}

	// 	log.Error().Msg("O certificado ICP deste usuário está expirado.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "GSS257", "O certificado ICP deste usuário está expirado."))
	// }

	return ctx.Next()
}
