package product

import (
	"encoding/json"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

type Product struct {
	Uid           string              `json:"uid,omitempty" redis:"uid,omitempty"`
	Name          string              `json:"name,omitempty" redis:"name,omitempty"`
	Description   string              `json:"description,omitempty" redis:"description,omitempty"`
	Pictures      []string            `json:"pictures,omitempty" redis:"pictures,omitempty"`
	CategoryCodes []string            `json:"-" redis:"category_codes,omitempty"`
	Categories    []category.Category `json:"categories,omitempty" redis:"-"`
	Price         string              `json:"price,omitempty" redis:"price,omitempty"`
	Promotion     string              `json:"promotion,omitempty" redis:"promotion,omitempty"`
	Code          string              `json:"code,omitempty" redis:"code,omitempty"`
	Weight        string              `json:"weight,omitempty" redis:"weight,omitempty"`
	Color         string              `json:"color,omitempty" redis:"color,omitempty"`
	Status        Status              `json:"status,omitempty" redis:"status,omitempty"`
	CreatedAt     string              `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt     string              `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
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

func (p *Product) NewUid() string {
	p.Uid = utils.Nonce()
	return p.Uid
}

func (p *Product) AddCategory(ctg category.Category) []category.Category {
	p.Categories = append(p.Categories, ctg)

	return p.Categories
}

func (p *Product) RemoveCategory(ctg category.Category) []category.Category {
	var ctgs []category.Category = []category.Category{}

	for _, category := range p.Categories {
		if category.Code == ctg.Code {
			continue
		}

		ctgs = append(ctgs, category)
	}

	return p.Categories
}

func (p *Product) FindCategory(ctg category.Category) category.Category {
	var result category.Category

	for _, category := range p.Categories {
		if category.Code == ctg.Code {
			result = category
			break
		}
	}

	return result
}

func (p *Product) ForEachCategory(fn func(ctg category.Category)) []category.Category {
	for _, category := range p.Categories {
		fn(category)
	}

	return p.Categories
}

func (p *Product) ForEachCategoryCodes(fn func(code string)) []string {
	for _, code := range p.CategoryCodes {
		fn(code)
	}

	return p.CategoryCodes
}

func (p *Product) ApplyCategoryCodes() []string {
	p.CategoryCodes = []string{}

	for _, category := range p.Categories {
		p.CategoryCodes = append(p.CategoryCodes, category.Code)
	}

	p.CategoryCodes = str.UniqueInSlice(p.CategoryCodes)

	return p.CategoryCodes
}

func (p *Product) SetUid(uid string) string {
	p.Uid = uid
	return p.Uid
}

func (p *Product) SetStatus(status Status) Status {
	p.Status = status
	return status
}

func (p *Product) SetName(name string) string {
	p.Name = name
	return p.Name
}

func (p *Product) SetDescription(description string) string {
	p.Description = description
	return p.Description
}

func (p *Product) SetCategoryCodes(codes []string) []string {
	p.CategoryCodes = codes
	return p.CategoryCodes
}

func (p *Product) SetPictures(pictures []string) []string {
	p.Pictures = pictures
	return p.Pictures
}

func (p *Product) SetPrice(price string) string {
	p.Price = price
	return p.Price
}

func (p *Product) SetPromotion(promotion string) string {
	p.Promotion = promotion
	return p.Promotion
}

func (p *Product) SetCode(code string) string {
	p.Code = code
	return p.Code
}

func (p *Product) SetColor(color string) string {
	p.Color = color
	return p.Color
}

func (p *Product) SetWeight(weight string) string {
	p.Weight = weight
	return p.Weight
}

func (p *Product) SetCreatedAt(createdAt string) string {
	p.CreatedAt = createdAt
	return p.CreatedAt
}

func (p *Product) SetUpdatedAt(updatedAt string) string {
	p.UpdatedAt = updatedAt
	return p.UpdatedAt
}

func (p *Product) ToString() (string, error) {
	var log = logger.New()

	productJson, err := json.Marshal(p)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear o struc para JSON. (%v)", p.Name)
		return "", err
	}
	return string(productJson), nil
}

func (p *Product) Inject(product *Product) *Product {

	if product.Name != "" {
		p.Name = product.Name
	}

	if product.Description != "" {
		p.Description = product.Description
	}

	if fmt.Sprintf("%T", product.CategoryCodes) == "[]string" {
		p.CategoryCodes = product.CategoryCodes
	}

	if fmt.Sprintf("%T", product.Pictures) == "[]string" {
		p.Pictures = product.Pictures
	}

	if product.Price != "" {
		p.Price = product.Price
	}

	if product.Promotion != "" {
		p.Promotion = product.Promotion
	}

	if product.Code != "" {
		p.Code = product.Code
	}

	if product.Color != "" {
		p.Color = product.Color
	}

	if product.Weight != "" {
		p.Weight = product.Weight
	}

	if product.Status != Undefined {
		p.Status = product.Status
	}

	if product.CreatedAt != "" {
		p.CreatedAt = product.CreatedAt
	}

	if product.UpdatedAt != "" {
		p.UpdatedAt = product.UpdatedAt
	}

	return p
}
