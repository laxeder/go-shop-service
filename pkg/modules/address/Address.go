package address

import "github.com/laxeder/go-shop-service/pkg/modules/redisdb"

type Address struct {
	Uuid         string `json:"uuid,omitempty" redis:"uuid,omitempty"`
	Uid          string `json:"uid,omitempty" redis:"uid,omitempty"`
	Number       string `json:"number,omitempty" redis:"number,omitempty"`
	Zip          string `json:"zip,omitempty" redis:"zip,omitempty"`
	Street       string `json:"street,omitempty" redis:"street,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty" redis:"neighborhood,omitempty"`
	City         string `json:"city,omitempty" redis:"city,omitempty"`
	State        string `json:"state,omitempty" redis:"state,omitempty"`
	redisdb.DataInfo
}
