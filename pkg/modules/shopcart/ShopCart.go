package shopcart

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/modules/freight"
	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
	"github.com/rs/zerolog/log"
)

type ShopCart struct {
	Uuid          string            `json:"uuid,omitempty" redis:"uuid,omitempty"`
	Products      []product.Product `json:"products,omitempty" redis:"-"`
	ProductsUid   []string          `json:"-" redis:"products_uid,omitempty"`
	ProductsTotal int               `json:"products_total,omitempty" redis:"-"`
	Items         int               `json:"items,omitempty" redis:"-"`
	Taxes         int               `json:"taxes,omitempty" redis:"taxes,omitempty"`
	Freights      []freight.Freight `json:"freights,omitempty" redis:"-"`
	FreightsUid   []string          `json:"-" redis:"freights_total,omitempty"`
	FreightsTotal int               `json:"freights_total,omitempty" redis:"-"`
	SubTotal      int               `json:"sub_total,omitempty" redis:"-"`
	Total         int               `json:"total,omitempty" redis:"-"`
	Status        Status            `json:"status,omitempty" redis:"status,omitempty"`
	LastAcesses   string            `json:"last_acesses,omitempty" redis:"last_acesses,omitempty"`
}

func New(userByte ...[]byte) (shopcart *ShopCart, err error) {
	shopcart = &ShopCart{}
	err = nil

	if len(userByte) == 0 {
		return
	}

	err = json.Unmarshal(userByte[0], shopcart)
	if err != nil {
		log.Error().Err(err).Msgf("O json do carrinho de compras está incorreto. %s", userByte[0])
		return
	}

	return
}

//? ************************ PRODUCTS ************************

func (s *ShopCart) Resume() {
	s.Items = len(s.Products)
	s.SubTotal = s.ProductsTotal + s.FreightsTotal
	s.Total = s.SubTotal + s.Taxes
}

func (s *ShopCart) Add(product product.Product) {
	s.Products = append(s.Products, product)

	s.ProductsTotal += product.Price
}

func (s *ShopCart) Lote(products []product.Product) {
	for _, product := range products {
		s.Add(product)
	}
}

func (s *ShopCart) Remove(prod product.Product) {
	var prods []product.Product

	for _, p := range s.Products {
		if p.Uid == prod.Uid {
			s.ProductsTotal -= prod.Price

			continue
		}

		prods = append(prods, prod)
	}

	s.Products = prods
}

func (s *ShopCart) RemoveLote(products []product.Product) {
	for _, p := range products {
		s.Remove(p)
	}
}

func (s *ShopCart) ApplyProductsUid() []string {
	s.ProductsUid = []string{}

	for _, product := range s.Products {
		s.ProductsUid = append(s.ProductsUid, product.Uid)
	}

	s.ProductsUid = str.UniqueInSlice(s.ProductsUid)

	return s.ProductsUid
}

func (s *ShopCart) ForEachProductsUid(fn func(uid string)) []string {
	for _, uid := range s.ProductsUid {
		fn(uid)
	}

	return s.ProductsUid
}

//? ************************ FREIGHT ************************

func (s *ShopCart) FreightResume() int {
	s.FreightsTotal = 0

	for _, f := range s.Freights {
		s.FreightsTotal += f.Price
	}

	return s.FreightsTotal
}

func (s *ShopCart) AddFreight(f freight.Freight) {
	s.Freights = append(s.Freights, f)

	s.FreightsTotal += f.Price

}

func (s *ShopCart) LoteFreights(freights []freight.Freight) {
	for _, f := range freights {
		s.AddFreight(f)
	}
}

func (s *ShopCart) RemoveFreight(freig freight.Freight) {
	var freigs []freight.Freight

	for _, f := range s.Freights {
		if f.Uid == freig.Uid {
			s.FreightsTotal -= f.Price

			continue
		}

		freigs = append(freigs, f)
	}

	s.Freights = freigs
}

func (s *ShopCart) RemoveLoteFreights(freights []freight.Freight) {
	for _, f := range freights {
		s.RemoveFreight(f)
	}
}

func (s *ShopCart) ApplyFreightUid() []string {
	s.FreightsUid = []string{}

	for _, freight := range s.Products {
		s.FreightsUid = append(s.FreightsUid, freight.Uid)
	}

	s.FreightsUid = str.UniqueInSlice(s.FreightsUid)

	return s.FreightsUid
}

func (s *ShopCart) ForEachFreightsUid(fn func(uid string)) []string {
	for _, uid := range s.FreightsUid {
		fn(uid)
	}

	return s.FreightsUid
}
