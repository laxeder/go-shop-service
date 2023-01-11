package utils

import (
	"crypto/sha512"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/utils/str"
	"github.com/laxeder/go-shop-service/pkg/utils/tokens"
)

func NewSalt() string {
	return tokens.Nonce()
}

func NewHashPassword(salt string, password string) (hash string) {
	hash = ""
	h := sha512.New()
	h.Write([]byte(str.MixStrings(salt, password)))
	hash = string(fmt.Sprintf("%x\n", h.Sum(nil)))

	return hash
}
