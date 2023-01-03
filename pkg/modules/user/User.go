package user

import (
	"encoding/json"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var log = logger.New()

type User struct {
	Uuid            string `json:"uuid,omitempty" redis:"uuid,omitempty"`
	Fullname        string `json:"full_name,omitempty" redis:"full_name,omitempty"`
	FirstName       string `json:"first_name,omitempty" redis:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty" redis:"last_name,omitempty"`
	Document        string `json:"document,omitempty" redis:"document,omitempty"`
	Email           string `json:"email,omitempty" redis:"email,omitempty"`
	Telephone       string `json:"telephone,omitempty" redis:"telephone,omitempty"`
	Password        string `json:"password,omitempty" redis:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty" redis:"-,omitempty"`
	Salt            string `json:"salt,omitempty" redis:"salt,omitempty"`
	Status          Status `json:"status,omitempty" redis:"status,omitempty"`
	CreatedAt       string `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func New(userByte ...[]byte) (user *User, err error) {
	user = &User{}
	err = nil

	if len(userByte) == 0 {
		return
	}

	err = json.Unmarshal(userByte[0], user)
	if err != nil {
		log.Error().Err(err).Msgf("O json do usuário está incorreto. %s", userByte[0])
		return
	}

	return
}

func (u *User) SetFullname() string {
	u.Fullname = fmt.Sprintf("%v %v", u.FirstName, u.LastName)
	return u.Fullname
}

func (u *User) SetDocument(document string) string {
	u.Document = str.DocumentPad(document)
	return u.Document
}

func (u *User) Inject(user *User) *User {

	if user.Uuid != "" {
		u.Uuid = user.Uuid
	}

	if user.Fullname != "" {
		u.Fullname = user.Fullname
	}

	if user.FirstName != "" {
		u.FirstName = user.FirstName
	}

	if user.LastName != "" {
		u.LastName = user.LastName
	}

	if user.Document != "" {
		u.Document = user.Document
	}

	if user.Email != "" {
		u.Email = user.Email
	}

	if user.Telephone != "" {
		u.Telephone = user.Telephone
	}

	if user.Password != "" {
		u.Password = user.Password
	}

	if user.ConfirmPassword != "" {
		u.ConfirmPassword = user.ConfirmPassword
	}

	if user.Salt != "" {
		u.Salt = user.Salt
	}

	if user.Status != Undefined {
		u.Status = user.Status
	}

	if user.CreatedAt != "" {
		u.CreatedAt = user.CreatedAt
	}

	if user.UpdatedAt != "" {
		u.UpdatedAt = user.UpdatedAt
	}

	return u
}

func (u *User) InjectMap(userMap any) *User {
	var log = logger.New()

	b, err := json.Marshal(userMap)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao transformar user map em byte %s", userMap)
		return u
	}

	err = json.Unmarshal(b, &u)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao transformar byte em user map %s", userMap)
		return u
	}

	return u
}
