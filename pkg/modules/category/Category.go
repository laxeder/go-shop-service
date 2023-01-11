package category

import (
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/utils/str"
)

type Category struct {
	Name        string `json:"name,omitempty" redis:"name,omitempty"`
	Description string `json:"description,omitempty" redis:"description,omitempty"`
	Code        string `json:"code,omitempty" redis:"code,omitempty"`
	redisdb.DataInfo
}

func (c *Category) GenerateCode() string {
	keys, err := Repository().GetKeys()

	if err != nil {
		c.Code = str.PadCategory("0")
		return c.Code
	}

	c.Code = str.PadCategory(fmt.Sprintf("%v", len(keys)))

	return c.Code
}
