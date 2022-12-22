package user

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

type UserLogin struct {
	Email    string `json:"email,omitempty" redis:"email,omitempty"`
	Password string `json:"password,omitempty" redis:"password,omitempty"`
	Salt     string `json:"salt,omitempty" redis:"salt,omitempty"`
}

func (u *UserLogin) HashPassword(salt, password string) (hash string) {
	hash = ""
	h := sha512.New()
	h.Write([]byte(str.MixStrings(salt, password)))
	hash = string(fmt.Sprintf("%x\n", h.Sum(nil)))
	return hash
}

func NewUserLogin(userLoginByte ...[]byte) (userLogin *UserLogin, err error) {
	var log = logger.New()

	userLogin = &UserLogin{}
	err = nil

	if len(userLoginByte) == 0 {
		return userLogin, err
	}

	err = json.Unmarshal(userLoginByte[0], userLogin)
	if err != nil {
		log.Error().Err(err).Msgf("O json do usuário está incorreto. %v", userLoginByte[0])
		return userLogin, err
	}

	return userLogin, err
}
