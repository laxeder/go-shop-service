package shopcart

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/laxeder/go-shop-service/pkg/modules/product"
	"github.com/rs/zerolog/log"
)

type ShopCart struct {
	Uuid               string                  `json:"uuid,omitempty" redis:"uuid,omitempty"`
	ProductsResume     []product.ProductResume `json:"products_resume,omitempty" redis:"products_resume,omitempty"`
	Products           []product.Product       `json:"products,omitempty" redis:"-"`
	ProductsUid        []string                `json:"-" redis:"products_uid,omitempty"`
	ProductsTotal      int                     `json:"products_total,omitempty" redis:"-"`
	Items              int                     `json:"items,omitempty" redis:"-"`
	Taxes              int                     `json:"taxes,omitempty" redis:"taxes,omitempty"`
	FreightsTotal      int                     `json:"freights_total,omitempty" redis:"-"`
	SubTotal           int                     `json:"sub_total,omitempty" redis:"-"`
	Total              int                     `json:"total,omitempty" redis:"-"`
	DiscountPercentage int                     `json:"discount_percentage,omitempty" redis:"discount_percentage,omitempty"`
	Status             Status                  `json:"status,omitempty" redis:"status,omitempty"`
	LastAcesses        string                  `json:"last_acesses,omitempty" redis:"last_acesses,omitempty"`
}

func New(shopcartByte ...[]byte) (shopcart *ShopCart, err error) {
	shopcart = &ShopCart{}
	err = nil

	if len(shopcartByte) == 0 {
		return
	}

	err = json.Unmarshal(shopcartByte[0], shopcart)
	if err != nil {
		log.Error().Err(err).Msgf("O json do carrinho de compras está incorreto. %s", shopcartByte[0])
		return
	}

	return
}

func (s *ShopCart) CalcPercentage(price int, percentage int) int {
	return price - int(math.Round(float64(price)*(float64(percentage)/100)))
}

//? ************************ PRODUCTS ************************

func (s *ShopCart) ReduceProducts(products []product.Product) []product.ProductResume {
	var productsResume []product.ProductResume

	for _, prod := range products {
		productResume := product.ProductResume{
			ProductUid:      prod.Uid,
			FreightUid:      prod.Freight.Uid,
			Product:         prod,
			Freight:         prod.Freight,
			ZipcodeReceiver: prod.Freight.ZipcodeReceiver,
		}

		productsResume = append(productsResume, productResume)
	}

	return productsResume
}

func (s *ShopCart) ExpandProducts(productsResume []product.ProductResume) []product.Product {
	var products []product.Product

	for _, productRes := range productsResume {
		productRes.Freight.ZipcodeReceiver = productRes.ZipcodeReceiver
		productRes.Product.Freight = productRes.Freight

		products = append(products, productRes.Product)
	}

	s.Products = products

	return products
}

func (s *ShopCart) Resume() {
	s.Items = len(s.Products)
	s.SubTotal = s.ProductsTotal + s.FreightsTotal
	s.Total = s.CalcPercentage(s.SubTotal+s.Taxes, s.DiscountPercentage)
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
	isRemoved := false

	for _, p := range s.Products {
		if !isRemoved && p.Uid == prod.Uid {
			s.ProductsTotal -= prod.Price
			isRemoved = true
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

	for _, prod := range s.Products {
		s.FreightsTotal += prod.Freight.Price
	}

	return s.FreightsTotal
}

func (s *ShopCart) Inject(shopcart *ShopCart) *ShopCart {

	if shopcart.Uuid != "" {
		s.Uuid = shopcart.Uuid
	}

	if fmt.Sprintf("%T", shopcart.Products) == "[]product.ProductResume" {
		s.ProductsResume = shopcart.ProductsResume
	}

	if fmt.Sprintf("%T", shopcart.Products) == "[]product.Product" {
		s.Products = shopcart.Products
	}

	if fmt.Sprintf("%T", shopcart.ProductsUid) == "[]string" {
		s.ProductsUid = shopcart.ProductsUid
	}

	if shopcart.ProductsTotal != 0 {
		s.ProductsTotal = shopcart.ProductsTotal
	}

	if shopcart.FreightsTotal != 0 {
		s.FreightsTotal = shopcart.FreightsTotal
	}

	if shopcart.Items != 0 {
		s.Items = shopcart.Items
	}

	if shopcart.Taxes != 0 {
		s.Taxes = shopcart.Taxes
	}

	if shopcart.SubTotal != 0 {
		s.SubTotal = shopcart.SubTotal
	}

	if shopcart.Total != 0 {
		s.Total = shopcart.Total
	}

	if shopcart.DiscountPercentage != 0 {
		s.DiscountPercentage = shopcart.DiscountPercentage
	}

	if shopcart.Status != "" {
		s.Status = shopcart.Status
	}

	if shopcart.LastAcesses != "" {
		s.LastAcesses = shopcart.LastAcesses
	}

	return s
}
