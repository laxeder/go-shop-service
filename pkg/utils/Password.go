package utils

import (
	"crypto/sha512"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

func NewSalt() string {
	return Nonce()
}

func NewHashPassword(salt string, password string) (hash string) {
	hash = ""
	h := sha512.New()
	h.Write([]byte(str.MixStrings(salt, password)))
	hash = string(fmt.Sprintf("%x\n", h.Sum(nil)))

	return hash
}
