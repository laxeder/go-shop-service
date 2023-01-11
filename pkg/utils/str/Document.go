package str

import (
	"fmt"
	"strings"

	"github.com/klassmann/cpfcnpj"
)

func DocumentClean(docuemnt string) string {
	return OnlyNumbers(cpfcnpj.Clean(docuemnt))
}

// pad start: 14 positions - 00000000000000
func DocumentPad(document string) string {
	pad := strings.Replace(fmt.Sprintf("%14s", DocumentClean(document)), " ", "X", 14)
	return pad[:14]
}

func DocumentValid(docuemnt string) bool {
	doc := OnlyNumbers(cpfcnpj.Clean(docuemnt))
	return cpfcnpj.ValidateCNPJ(doc) || cpfcnpj.ValidateCPF(doc)
}
