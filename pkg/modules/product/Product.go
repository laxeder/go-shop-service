package product

import (
	"encoding/json"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type Product struct {
	Name        string   `json:"name,omitempty" redis:"name,omitempty"`
	Description string   `json:"description,omitempty" redis:"description,omitempty"`
	Pictures    []string `json:"pictures,omitempty" redis:"pictures,omitempty"`
	Categorys   []string `json:"categorys,omitempty" redis:"categorys,omitempty"`
	Price       string   `json:"price,omitempty" redis:"price,omitempty"`
	Promotion   string   `json:"promotion,omitempty" redis:"promotion,omitempty"`
	Code        string   `json:"code,omitempty" redis:"code,omitempty"`
	Weight      string   `json:"weight,omitempty" redis:"weight,omitempty"`
	Color       string   `json:"color,omitempty" redis:"color,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt   string   `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func New(productByte ...[]byte) (product *Product, err error) {
	var log = logger.New()

	product = &Product{}
	err = nil

	if len(productByte) == 0 {
		return product, err
	}

	err = json.Unmarshal(productByte[0], product)
	if err != nil {
		log.Error().Err(err).Msgf("O json do product está incorreto. %v", productByte[0])
		return product, err
	}

	return product, err
}

func (a *Product) SetName(name string) string {
	a.Name = name
	return a.Name
}

func (a *Product) SetDescription(description string) string {
	a.Description = description
	return a.Description
}

func (a *Product) SetCategorys(categorys []string) []string {
	a.Categorys = categorys
	return a.Categorys
}

func (a *Product) SetPictures(pictures []string) []string {
	a.Pictures = pictures
	return a.Pictures
}

func (a *Product) SetPrice(price string) string {
	a.Price = price
	return a.Price
}

func (a *Product) SetPromotion(promotion string) string {
	a.Promotion = promotion
	return a.Promotion
}

func (a *Product) SetCode(code string) string {
	a.Code = code
	return a.Code
}

func (a *Product) SetColor(color string) string {
	a.Color = color
	return a.Color
}

func (a *Product) SetWeight(weight string) string {
	a.Weight = weight
	return a.Weight
}

func (a *Product) SetCreatedAt(createdAt string) string {
	a.CreatedAt = createdAt
	return a.CreatedAt
}

func (a *Product) SetUpdatedAt(updatedAt string) string {
	a.UpdatedAt = updatedAt
	return a.UpdatedAt
}

func (a *Product) ToString() (string, error) {
	var log = logger.New()

	productJson, err := json.Marshal(a)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear a struc para JSON. (%v)", a.Name)
		return "", err
	}
	return string(productJson), nil
}

func (a *Product) Inject(product *Product) *Product {

	if product.Name != "" {
		a.Name = product.Name
	}

	if product.Description != "" {
		a.Description = product.Description
	}

	if fmt.Sprintf(" %T", product.Categorys) != "[]string" {
		a.Categorys = product.Categorys
	}

	if fmt.Sprintf(" %T", product.Pictures) != "[]string" {
		a.Pictures = product.Pictures
	}

	if product.Price != "" {
		a.Price = product.Price
	}

	if product.Promotion != "" {
		a.Promotion = product.Promotion
	}

	if product.Code != "" {
		a.Code = product.Code
	}

	if product.Color != "" {
		a.Color = product.Color
	}

	if product.Weight != "" {
		a.Weight = product.Weight
	}

	if product.CreatedAt != "" {
		a.CreatedAt = product.CreatedAt
	}

	if product.UpdatedAt != "" {
		a.UpdatedAt = product.UpdatedAt
	}

	return a
}
