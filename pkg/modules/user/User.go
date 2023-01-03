package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

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
}

func (u *User) GenerateUuid() {
	u.Uuid = uuid.New().String()
}

func (u *User) GenerateFullName() {
	u.Fullname = fmt.Sprintf("%v %v", u.FirstName, u.LastName)
}

func (u *User) GenerateDocument() {
	u.Document = str.DocumentPad(u.Document)
}

func (u *User) GeneratePassword() {
	u.Salt = utils.NewSalt()
	u.Password = utils.NewHashPassword(u.Salt, u.Password)
}
