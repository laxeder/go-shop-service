package product

import (
	"fmt"

	"github.com/ledongthuc/goterators"
)

func ConvertCategories(ctg []Categories) []string {

	strs := goterators.Map(ctg, func(item Categories) string {
		str := fmt.Sprintf("%v", item)

		return string(str)
	})

	return strs
}
