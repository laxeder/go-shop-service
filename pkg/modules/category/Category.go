package category

import (
	"encoding/json"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

type Category struct {
	Name        string `json:"name,omitempty" redis:"name,omitempty"`
	Description string `json:"description,omitempty" redis:"description,omitempty"`
	Code        string `json:"code,omitempty" redis:"code,omitempty"`
	Status      Status `json:"status,omitempty" redis:"status,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty" redis:"created_at,omitempty"`
}

func New(categoryByte ...[]byte) (category *Category, err error) {
	var log = logger.New()

	category = &Category{}
	err = nil

	if len(categoryByte) == 0 {
		return
	}

	err = json.Unmarshal(categoryByte[0], category)
	if err != nil {
		log.Error().Err(err).Msgf("O json da categoria está incorreto. %s", categoryByte[0])
		return
	}

	return
}

func (c *Category) NewCode() string {
	keys, err := Repository().GetKeys()

	if err != nil {
		c.Code = str.PadCategory("0")
		return c.Code
	}

	c.Code = str.PadCategory(fmt.Sprintf("%v", len(keys)))
	return c.Code
}

func (c *Category) ToString() (string, error) {
	var log = logger.New()

	categoryJson, err := json.Marshal(c)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear o struc para JSON. (%v)", c.Name)
		return "", err
	}
	return string(categoryJson), nil
}

func (c *Category) Inject(category *Category) *Category {

	if category.Name != "" {
		c.Name = category.Name
	}

	if category.Description != "" {
		c.Description = category.Description
	}

	if category.Code != "" {
		c.Code = category.Code
	}

	if category.Status != "" {
		c.Status = category.Status
	}

	if category.UpdatedAt != "" {
		c.UpdatedAt = category.UpdatedAt
	}

	if category.CreatedAt != "" {
		c.CreatedAt = category.CreatedAt
	}

	return c
}
