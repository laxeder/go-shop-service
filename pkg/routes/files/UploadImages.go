package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func UploadImages(ctx *fiber.Ctx) error {
	// var log = logger.New()

	// body := ctx.Body()
	// document := ctx.Params("document")

	// transforma o json em  struct
	// imagesBody, err := images.New(body)
	// if err != nil {
	// 	log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "XXXX1", "O formado dos dados envidados está incorreto."))
	// }

	// for i, image := range imagesBody.Images {

	// 	log.Print(regex.MimeType.FindString(image.Content))
	// 	if !str.ValidRegex(regex.IsBase64, image.Content) {
	// 		return response.Ctx(ctx).Result(response.Error(400, "XXXX2", "O formado dos dados envidados está incorreto."))
	// 	}

	// 	if checkSizeImage(image.Content) {
	// 		return response.Ctx(ctx).Result(response.Error(400, "XXXX3", "O tamanho máximo permitido é de 1 mb"))
	// 	}

	// 	imagesBody.Images[i].UUID = images.Repository().NewUuid()
	// 	imagesBody.Images[i].Size = fmt.Sprintf("%d", calculateSize(image.Content))
	// 	imagesBody.Images[i].MimeType = regex.MimeType.FindString(image.Content)

	// }

	// imagesBody.Document = str.DocumentPad(document)

	// dataUploadImages, err := images.Repository().GetImageByDocument(document)
	// if err != nil {
	// 	log.Error().Err(err).Msg("O formado dos dados envidados está incorreto.")
	// 	return response.Ctx(ctx).Result(response.Error(400, "XXXX4", "O formado dos dados envidados está incorreto."))
	// }

	// if len(dataUploadImages.UUID) == 0 {
	// 	imagesBody.NewUuid()
	// 	imagesBody.CreatedAt = date.NowUTC()
	// 	imagesBody.UpdatedAt = date.NowUTC()
	// 	images.Repository().Save(imagesBody)

	// 	return response.Ctx(ctx).Result(response.Success(201))
	// }

	// dataUploadImages.UpdatedAt = date.NowUTC()
	// dataUploadImages.Images = imagesBody.Images
	// images.Repository().Save(dataUploadImages)

	return response.Ctx(ctx).Result(response.Success(201))
}
