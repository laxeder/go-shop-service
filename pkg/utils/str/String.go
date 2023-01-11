package str

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/laxeder/go-shop-service/pkg/utils/regex"
)

func String(str any) string {
	return fmt.Sprintf("%v", reflect.ValueOf(str))
}

func Dump(stc any) {
	fmt.Printf("%#+v", stc)
}

func Substring(input string, start int, end int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+end > len(asRunes) {
		end = len(asRunes) - start
	}

	return string(asRunes[start : start+end])
}

func Regex(regex *regexp.Regexp, value string) bool {
	return regex.MatchString(value)
}

func OnlyNumbers(str string) string {
	return regex.OnlyNumber.ReplaceAllString(str, "")
}
