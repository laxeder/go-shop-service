package user

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
	"github.com/laxeder/go-shop-service/pkg/utils"
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
		log.Error().Err(err).Msgf("O json do usuário está incorreto. %v", userByte[0])
		return
	}

	return
}

func (u *User) SetFullname() string {
	u.Fullname = fmt.Sprintf("%v %v", u.FirstName, u.LastName)
	return u.Fullname
}

func (u *User) SetFirstName(firstName string) string {
	u.FirstName = firstName
	return u.FirstName
}

func (u *User) SetLastName(lastName string) string {
	u.LastName = lastName
	return u.LastName
}

func (u *User) SetDocument(document string) string {
	u.Document = str.DocumentPad(document)
	return u.Document
}

func (u *User) SetEmail(email string) string {
	u.Email = email
	return u.Email
}

func (u *User) SetTelephone(telephone string) string {
	u.Telephone = telephone
	return u.Telephone
}

func (u *User) SetPassword(password string) string {
	u.Password = password
	return u.Password
}

func (u *User) SetConfirmPassword(confirmPassword string) string {
	u.ConfirmPassword = confirmPassword
	return u.ConfirmPassword
}

func (u *User) SetSalt(salt string) string {
	u.Salt = salt
	return u.Salt
}

func (u *User) SetStatus(status Status) Status {
	u.Status = status
	return status
}

func (u *User) NewSalt() string {
	u.Salt = utils.Nonce()
	return u.Salt
}

func (u *User) NewHashPassword() string {
	u.Password = u.HashPassword(u.Salt, u.Password)
	return u.Password
}

func (u *User) HashPassword(salt, password string) (hash string) {
	hash = ""
	h := sha512.New()
	h.Write([]byte(str.MixStrings(salt, password)))
	hash = string(fmt.Sprintf("%x\n", h.Sum(nil)))
	return hash
}

func (a *User) NewUuid() string {
	a.Uuid = uuid.New().String()
	return a.Uuid
}

func (a *User) SetUuid(uuid string) string {
	a.Uuid = uuid
	return a.Uuid
}

func (a *User) SetCreatedAt(createdAt string) string {
	a.CreatedAt = createdAt
	return a.CreatedAt
}

func (a *User) SetUpdatedAt(updatedAt string) string {
	a.UpdatedAt = updatedAt
	return a.UpdatedAt
}

func (u *User) ToString() (userString string, err error) {

	var log = logger.New()

	userString = ""
	err = nil

	userJson, err := json.Marshal(u)
	if err != nil {
		log.Error().Err(err).Msgf("A struct do usuário está incorreta. %v", u.Document)
		return
	}
	userString = string(userJson)
	return
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
