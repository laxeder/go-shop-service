package auth

import (
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

type Login struct {
	Document string `json:"document"`
	Password string `json:"password"`
}

func LogIn(loginByte ...[]byte) (login *Login, err error) {

	// var log = logger.New()

	login = &Login{}
	err = nil

	// if len(loginByte) == 0 {
	// 	return
	// }

	// err = json.Unmarshal(loginByte[0], login)
	// if err != nil {
	// 	log.Error().Err(err).Msgf("O json do usuário está incorreto. %v", loginByte[0])
	// 	return
	// }

	return
}

func (l *Login) Valid() (result *response.Result) {
	user := &user.User{Document: l.Document, Password: l.Password}

	result = &response.Result{}

	result = user.DocumentValid()
	if result.Status != 200 {
		return
	}

	result = user.PasswordValid()
	if result.Status != 200 {
		return
	}

	result = response.Success(200)
	return
}
