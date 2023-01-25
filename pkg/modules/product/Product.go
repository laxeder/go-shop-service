package product

import (
	"github.com/laxeder/go-shop-service/pkg/modules/category"
	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/utils/str"
	"github.com/laxeder/go-shop-service/pkg/utils/tokens"
)

type Product struct {
	Uid         string              `json:"uid,omitempty" redis:"uid,omitempty"`
	Name        string              `json:"name,omitempty" redis:"name,omitempty"`
	Description string              `json:"description,omitempty" redis:"description,omitempty"`
	Pictures    []string            `json:"pictures,omitempty" redis:"pictures,omitempty"`
	Categories  []category.Category `json:"categories,omitempty" redis:"-"`
	Freights    []freight.Freight   `json:"freights,omitempty" redis:"-"` //? Lista de fretes dísponivel para esse produto
	Color       string              `json:"color,omitempty" redis:"color,omitempty"`
	Promotion   int                 `json:"promotion,omitempty" redis:"promotion,omitempty"` //TODO: Criar uma classe de promoções
	Code        string              `json:"code,omitempty" redis:"code,omitempty"`
	Price       int                 `json:"price,omitempty" redis:"price,omitempty"`
	Weight      int                 `json:"weight,omitempty" redis:"weight,omitempty"`
	redisdb.DataInfo
}

func (p *Product) Generate() {
	p.Uid = tokens.Nonce()
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

func (p *Product) GetCategoryCodes() []string {
	categoryCodes := []string{}

	for _, category := range p.Categories {
		categoryCodes = append(categoryCodes, category.Code)
	}

	categoryCodes = str.UniqueInSlice(categoryCodes)

	return categoryCodes
}

func (p *Product) GetFreightsUid() []string {
	freightsUid := []string{}

	for _, freight := range p.Freights {
		freightsUid = append(freightsUid, freight.Uid)
	}

	freightsUid = str.UniqueInSlice(freightsUid)

	return freightsUid
}
