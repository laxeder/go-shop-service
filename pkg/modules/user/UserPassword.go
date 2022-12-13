package user

import (
	"encoding/json"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
)

type UserPassword struct {
	Password           string `json:"password,omitempty" redis:"password,omitempty"`
	NewConfirmPassword string `json:"new_confirm_password,omitempty" redis:"new_confirm_password,omitempty"`
	NewPassword        string `json:"new_password,omitempty" redis:"new_password,omitempty"`
}

func NewUserPassowrd(userPassowordByte ...[]byte) (userPassword *UserPassword, err error) {
	var log = logger.New()

	userPassword = &UserPassword{}
	err = nil

	if len(userPassowordByte) == 0 {
		return userPassword, err
	}

	err = json.Unmarshal(userPassowordByte[0], userPassword)
	if err != nil {
		log.Error().Err(err).Msgf("O json do usuário está incorreto. %v", userPassowordByte[0])
		return userPassword, err
	}

	return userPassword, err
}
