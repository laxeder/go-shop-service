package user

import (
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func (u *User) Valid() *response.Result {

	checkFirstname := u.FirstNameValid()
	if checkFirstname.Status != 200 {
		return checkFirstname
	}

	checkLastName := u.LastNameValid()
	if checkLastName.Status != 200 {
		return checkLastName
	}

	checkDocument := u.DocumentValid()
	if checkDocument.Status != 200 {
		return checkDocument
	}

	checkEmail := u.EmailValid()
	if checkEmail.Status != 200 {
		return checkEmail
	}

	checkTelefone := u.TelephoneValid()
	if checkTelefone.Status != 200 {
		return checkTelefone
	}

	checkPassword := u.PasswordValid()
	if checkPassword.Status != 200 {
		return checkPassword
	}

	checkConfirmPassword := u.ConfirmPasswordValid()
	if checkConfirmPassword.Status != 200 {
		return checkConfirmPassword
	}

	return response.Success(200)
}

func (u *User) FirstNameValid() *response.Result {
	// return shared.NameValid(u.FirstName)
	return response.Success(200)
}

func (u *User) LastNameValid() *response.Result {
	// return shared.NameValid(u.LastName)
	return response.Success(200)
}

func (u *User) DocumentValid() *response.Result {
	// return shared.DocumentValid(u.Document)
	return response.Success(200)
}

func (u *User) EmailValid() *response.Result {
	// return shared.EmailValid(u.Email)
	return response.Success(200)

}

func (u *User) TelephoneValid() *response.Result {

	// var log = logger.New()

	// if u.Telephone == "" {
	// 	log.Error().Msgf("O campo do email não pode ser vazio. (%v)", u.Document)
	// 	return response.Error(400, "GSS018", "O campo telefone não pode ser vazio.")
	// }

	// if len(u.Telephone) < 11 {
	// 	log.Error().Msgf("O telefone não poder ser menor que 11 números. Favor manter o fomrmato (XX) XXXXX-XXX. (%v)", u.Document)
	// 	return response.Error(400, "GSS019", "O telefone não poder ser menor que 11 números. Favor manter o fomrmato (XX) XXXXX-XXX.")
	// }

	// if len(u.Telephone) > 15 {
	// 	log.Error().Msgf("O telefone não poder ser maior que 11 números. Favor manter o fomrmato (XX) XXXXX-XXX. (%v)", u.Document)
	// 	return response.Error(400, "GSS020", "O telefone não poder ser maior que 11 números. Favor manter o fomrmato (XX) XXXXX-XXX.")
	// }

	// if regex.HasLetter.MatchString(u.Telephone) {
	// 	log.Error().Msgf("O telefone não poder pode conter letras. (%v)", u.Document)
	// 	return response.Error(400, "GSS021", "O telefone não poder pode conter letras.")
	// }

	// if regex.HasCharSpecialsToPhone.MatchString(u.Telephone) {
	// 	log.Error().Msgf("O telefone não poder pode conter letras especiais ou acentuação. (%v)", u.Document)
	// 	return response.Error(400, "GSS022", "O telefone não poder pode conter letras especiais ou acentuação.")
	// }

	return response.Success(200)
}

func (u *User) PasswordValid() *response.Result {
	// return shared.PasswordValid(u.Password, u.Document)
	return response.Success(200)
}

func (u *User) ConfirmPasswordValid() *response.Result {

	// var log = logger.New()

	// if u.ConfirmPassword == "" {
	// 	log.Error().Msgf("O campo confirmar senha não pode ser vazio. (%v)", u.Document)
	// 	return response.Error(400, "GSS029", "O campo confirmar senha não pode ser vazio.")
	// }

	// if u.Password != u.ConfirmPassword {
	// 	log.Error().Msgf("O campo confirmar senha deve ser idêntico ao da senha. (%v)", u.Document)
	// 	return response.Error(400, "GSS030", "O campo confirmar senha deve ser idêntico ao da senha.")
	// }

	return response.Success(200)
}
