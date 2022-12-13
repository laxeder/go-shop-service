package shared

import (
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/response"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

func UserCheck(document string) (userData *user.User, result *response.Result) {

	var log = logger.New()

	userData = &user.User{}
	result = &response.Result{}

	//!##################################################################################################################//
	//! VERIFICA SE O DOCUMENTO DO USUÁRIO EXISTE NA BASE DE DADOS 																				 							 //
	//!##################################################################################################################//
	userData, err := user.Repository().GetByDocument(document)
	if err != nil {
		log.Error().Err(err).Msgf("Os campos enviados estão incorretos. (%v)", document)
		result = response.ErrorDefault("BLC258")
		return
	}

	// verifica se o documento existe
	if len(userData.Document) == 0 {
		log.Error().Msgf("Este documento não existe na nossa base de dados. (%v)", document)
		result = response.Error(400, "BLC259", "Este documento não existe na nossa base de dados.")
		return
	}

	// verifica se a conta está desabilitada
	if userData.Status == user.Disabled {
		log.Error().Msgf("Esta conta está desabilitada por tempo indeterminado. (%v)", document)
		userData = &user.User{Document: document, Status: user.Disabled}
		result = response.Error(400, "BLC260", "Esta conta está desabilitada por tempo indeterminado.")
		return
	}

	return
}
