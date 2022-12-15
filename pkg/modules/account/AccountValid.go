package account

import (
	"github.com/laxeder/go-shop-service/pkg/modules/response"
)

func (a *Account) Valid() *response.Result {

	checkNickname := a.NicknameValid()
	if checkNickname.Status != 200 {
		return checkNickname
	}

	// checkProfession := a.ProfessionValid()
	// if checkProfession.Status != 200 {
	// 	return checkProfession
	// }

	checkRG := a.RGValid()
	if checkRG.Status != 200 {
		return checkRG
	}

	checkGender := a.GenderValid()
	if checkGender.Status != 200 {
		return checkGender
	}

	checkBirthday := a.BirthdayValid()
	if checkBirthday.Status != 200 {
		return checkBirthday
	}

	return response.Success(200)
}

func (a *Account) NicknameValid() *response.Result {
	// var log = logger.New()

	// 	if a.Nickname == "" {
	// 		log.Error().Msgf("O campo do apelido não pode ser vazio. (%v)", a.Document)
	// 		return response.Error(400, "BLC168", "O campo do apelido não pode ser vazio.")
	// 	}

	// 	if len(a.Nickname) <= 2 {
	// 		log.Error().Msgf("O apelido não poder ser menor que 2 caracteres. (%v)", a.Document)
	// 		return response.Error(400, "BLC169", "O apelido não poder ser menor que 2 caracteres.")
	// 	}

	// 	return response.Success(200)
	// }

	// func (a *Account) ProfessionValid() *response.Result {
	// 	var log = logger.New()

	// 	if a.Profession == "" {
	// 		log.Error().Msgf("O campo de profissão não pode ser vazio. (%v)", a.Document)
	// 		return response.Error(400, "BLC170", "O campo de profissão não pode ser vazio.")
	// 	}

	// 	if len(a.Profession) <= 2 {
	// 		log.Error().Msgf("O apelido não poder ser menor que 2 caracteres. (%v)", a.Document)
	// 		return response.Error(400, "BLC171", "O apelido não poder ser menor que 2 caracteres.")
	// 	}

	return response.Success(200)
}

func (a *Account) RGValid() *response.Result {
	// var log = logger.New()

	// if a.RG == "" {
	// 	log.Error().Msgf("O campo de RG não pode ser vazio. (%v)", a.Document)
	// 	return response.Error(400, "BLC172", "O campo de RG não pode ser vazio.")
	// }

	// if len(a.RG) <= 6 {
	// 	log.Error().Msgf("O campo de RG não poder ser menor que 6 caracteres. (%v)", a.Document)
	// 	return response.Error(400, "BLC173", "O campo de RG não poder ser menor que 6 caracteres.")
	// }

	// if regex.HasLetter.MatchString(a.RG) {
	// 	log.Error().Msgf("O RG não pode conter letras. (%v)", a.Document)
	// 	return response.Error(400, "BLC174", "O RG não pode conter letras.")
	// }

	return response.Success(200)
}

func (a *Account) GenderValid() *response.Result {
	// var log = logger.New()

	// if a.Gender == "" {
	// 	log.Error().Msgf("O campo de gênero não pode ser vazio. (%v)", a.Document)
	// 	return response.Error(400, "BLC175", "O campo de gênero não pode ser vazio.")
	// }

	// if string(a.Gender) != string(Male) && string(a.Gender) != string(Female) {
	// 	log.Error().Msgf("O opcao de sexo não está definida corretamente. (%v)", a.Document)
	// 	return response.Error(400, "BLC176", "O opcao de sexo não está definida corretamente.")
	// }

	return response.Success(200)
}

func (a *Account) BirthdayValid() *response.Result {
	// var log = logger.New()

	// if a.Birthday == "" {
	// 	log.Error().Msgf("O campo de data de nascimento não pode ser vazio. (%v)", a.Document)
	// 	return response.Error(400, "BLC177", "O campo de data de nascimento não pode ser vazio.")
	// }

	// if !date.BRValid(a.Birthday) {
	// 	log.Error().Msgf("O campo de data de nascimento não é uma data válida. (%v)", a.Document)
	// 	return response.Error(400, "BLC178", "O campo de data de nascimento não é uma data válida.")
	// }

	// if time.Now().Sub(date.BRToTime(a.Birthday)).Hours() < 0 {
	// 	log.Error().Msgf("O campo de data nascimento não pode estar no futuro. (%v)", a.Document)
	// 	return response.Error(400, "BLC179", "O campo de data nascimento não pode estar no futuro.")
	// }

	// var years float64 = 18
	// var hoursIn18years float64 = (((24 * 365) + 6) * years)

	// if time.Now().Sub(date.BRToTime(a.Birthday)).Hours() < hoursIn18years {
	// 	log.Error().Msgf("O campo de data nascimento não pode ser de um menor de idade. (%v)", a.Document)
	// 	return response.Error(400, "BLC180", "O campo de data nascimento não pode ser de um menor de idade.")
	// }

	// var yearsOld float64 = 120
	// var hoursIn120years float64 = (((24 * 365) + 6) * yearsOld)

	// if time.Now().Sub(date.BRToTime(a.Birthday)).Hours() > hoursIn120years {
	// 	log.Error().Msgf("O campo de data de nascimento não pode ser maior que 120 anos. (%v)", a.Document)
	// 	return response.Error(400, "BLC181", "O campo de data de nascimento não pode ser maior que 120 anos.")
	// }

	return response.Success(200)
}
