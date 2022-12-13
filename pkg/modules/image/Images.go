package image

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type Images struct {
	Document   string  `json:"document,omitempty" redis:"document,omitempty"`
	Images     []Image `json:"images,omitempty" redis:"images,omitempty"`
	ImagesJson string  `json:"-" redis:"images_json,omitempty"`
	Status     Status  `json:"status,omitempty" redis:"status,omitempty"`
}

func NewImages(imagesByte ...[]byte) (images *Images, err error) {

	var log = logger.New()

	images = &Images{}
	err = nil

	if len(imagesByte) == 0 {
		return
	}

	err = json.Unmarshal(imagesByte[0], images)
	if err != nil {
		log.Error().Err(err).Msgf("O json da imagem está incorreto. %v", imagesByte[0])
		return
	}

	return
}

func (i *Images) ImagesMerge() []Image {
	var imagesPool []Image

	if i.ImagesJson == "" {
		i.ImagesJson = "[]"
	}

	var imagesTemporary []Image
	if i.ImagesJson != "[]" {
		json.Unmarshal([]byte(i.ImagesJson), &imagesTemporary)
		for _, imageTemporary := range imagesTemporary {
			imagesPool = append(imagesPool, imageTemporary)
		}
	}

	if len(i.Images) > 0 {
		for _, image := range i.Images {
			imagesPool = append(imagesPool, image)
		}
	}

	i.Images = UniqueImages(imagesPool)
	i.ImagesJson = ImagesToJson(i.Images)

	return i.Images
}

func UniqueImages(images []Image) []Image {
	inResult := make(map[string]bool)
	var result []Image
	for _, image := range images {
		if _, ok := inResult[image.Uuid]; !ok {
			inResult[image.Uuid] = true
			result = append(result, image)
		}
		result = append(result, image)
	}
	return result
}

func ImagesToJson(images []Image) string {
	if len(images) == 0 {
		return "[]"
	}

	imageJson, err := json.Marshal(images)
	if err != nil {
		return "[]"
	}

	return string(imageJson)
}
