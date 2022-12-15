package utils

import (
	"fmt"
	"strings"

	"github.com/ledongthuc/goterators"
)

func ConvertToArrayString(matrix []any) []string {

	strs := goterators.Map(matrix, func(item any) string {
		str := fmt.Sprintf("%v", item)

		return string(str)
	})
	fmt.Println("Foos: " + strings.Join(strs, ","))

	return strs
}
