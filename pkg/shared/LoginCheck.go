package shared

import (
	"github.com/laxeder/go-shop-service/pkg/modules/auth"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func LoginCheck(login *auth.Login) (userData *user.User, result *response.Result) {

	var log = logger.New()

	userData = &user.User{}
	result = &response.Result{}

	// valida os campos de entrada
	result = login.Valid()
	if result.Status != 200 {
		return
	}

	userData, result = UserCheck(login.Document)
	if result.Status != 200 {
		return
	}

	user := user.User{Password: login.Password, Salt: userData.Salt}
	user.NewHashPassword()

	if user.Password != userData.Password {
		log.Error().Msgf("A senha está incorreta. Favor tentar outra senha. (%v)", login.Document)
		result = response.Error(400, "GSS261", "A senha está incorreta. Favor tentar outra senha.")
		return
	}

	return

}
