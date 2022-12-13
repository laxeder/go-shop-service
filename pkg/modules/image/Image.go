package image

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type Image struct {
	Uuid      string `json:"uuid,omitempty" redis:"uuid,omitempty"`
	Document  string `json:"document,omitempty" redis:"document,omitempty"`
	Size      string `json:"size,omitempty" redis:"size,omitempty"`
	MimeType  string `json:"mime_type,omitempty" redis:"mime_type,omitempty"`
	Name      string `json:"name,omitempty" redis:"name,omitempty"`
	Base64    string `json:"base64,omitempty" redis:"base64,omitempty"`
	Status    Status `json:"status,omitempty" redis:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func New(imageByte ...[]byte) (image *Image, err error) {

	var log = logger.New()

	image = &Image{}
	err = nil

	if len(imageByte) == 0 {
		return
	}

	err = json.Unmarshal(imageByte[0], image)
	if err != nil {
		log.Error().Err(err).Msgf("O json da imagem está incorreto. %v", imageByte[0])
		return
	}

	return
}

func (i *Image) NewUuid() string {
	i.Uuid = uuid.New().String()
	return i.Uuid
}

func (i *Image) ToString() (string, error) {
	var log = logger.New()

	imageJson, err := json.Marshal(i)
	if err != nil {
		log.Error().Err(err).Msgf("A struct da imagem está incorreta. %v", i.Document)
		return "", err
	}
	return string(imageJson), nil
}

func (i *Image) Inject(image *Image) *Image {

	if image.Uuid != "" {
		i.Uuid = image.Uuid
	}

	if image.Document != "" {
		i.Document = image.Document
	}

	if image.Size != "" {
		i.Size = image.Size
	}

	if image.MimeType != "" {
		i.MimeType = image.MimeType
	}

	if image.Name != "" {
		i.Name = image.Name
	}

	if image.Base64 != "" {
		i.Base64 = image.Base64
	}

	if image.CreatedAt != "" {
		i.CreatedAt = image.CreatedAt
	}

	if image.UpdatedAt != "" {
		i.UpdatedAt = image.UpdatedAt
	}

	return i
}
