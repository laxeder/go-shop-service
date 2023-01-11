package account

import (
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
)

type Account struct {
	Uuid       string `json:"uuid,omitempty" redis:"uuid,omitempty"`
	Nickname   string `json:"nickname,omitempty" redis:"nickname,omitempty"`
	Profession string `json:"profession,omitempty" redis:"profession,omitempty"`
	RG         string `json:"rg,omitempty" redis:"rg,omitempty"`
	Gender     Gender `json:"gender,omitempty" redis:"gender,omitempty"`
	Birthday   string `json:"birthday,omitempty" redis:"birthday,omitempty"`
	redisdb.DataInfo
}
