package account

import (
	"encoding/json"
	"fmt"

	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

type Account struct {
	Uid        string `json:"uid,omitempty" redis:"uid,omitempty"`
	Uuid       string `json:"uuid,omitempty" redis:"uuid,omitempty"`
	Nickname   string `json:"nickname,omitempty" redis:"nickname,omitempty"`
	Profession string `json:"profession,omitempty" redis:"profession,omitempty"`
	RG         string `json:"rg,omitempty" redis:"rg,omitempty"`
	Gender     Gender `json:"gender,omitempty" redis:"gender,omitempty"`
	Birthday   string `json:"birthday,omitempty" redis:"birthday,omitempty"`
	Options    bool   `json:"options,omitempty" redis:"options"`
	Status     Status `json:"status,omitempty" redis:"status,omitempty"`
	CreatedAt  string `json:"created_at,omitempty" redis:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty" redis:"updated_at,omitempty"`
}

func New(accountByte ...[]byte) (account *Account, err error) {
	var log = logger.New()

	account = &Account{}
	err = nil

	if len(accountByte) == 0 {
		return
	}

	err = json.Unmarshal(accountByte[0], account)
	if err != nil {
		log.Error().Err(err).Msgf("O json do account está incorreto. %v", accountByte[0])
		return
	}

	return
}

func (a *Account) NewUid() string {
	a.Uid = utils.Nonce()
	return a.Uid
}

func (a *Account) SetUid(uid string) string {
	a.Uid = uid
	return a.Uid
}
func (a *Account) SetUuid(uuid string) string {
	a.Uuid = uuid
	return a.Uuid
}

func (a *Account) SetNickname(nickname string) string {
	a.Nickname = nickname
	return a.Nickname
}

func (a *Account) SetProfession(profession string) string {
	a.Profession = profession
	return a.Profession
}

func (a *Account) SetRG(rg string) string {
	a.RG = rg
	return a.RG
}

func (a *Account) SetGender(gender Gender) Gender {
	a.Gender = gender
	return a.Gender
}

func (a *Account) SetBirthday(birthday string) string {
	a.Birthday = birthday
	return a.Birthday
}

func (a *Account) SetOptions(options bool) bool {
	a.Options = options
	return a.Options
}

func (a *Account) SetStatus(status Status) Status {
	a.Status = status
	return status
}

func (a *Account) SetCreatedAt(createdAt string) string {
	a.CreatedAt = createdAt
	return a.CreatedAt
}

func (a *Account) SetUpdatedAt(updatedAt string) string {
	a.UpdatedAt = updatedAt
	return a.UpdatedAt
}

func (a *Account) ToString() (string, error) {
	var log = logger.New()

	accountJson, err := json.Marshal(a)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear a struc para JSON. (%v)", a.Uid)
		return "", err
	}
	return string(accountJson), nil
}

func (a *Account) Inject(account *Account) *Account {

	if account.Uid != "" {
		a.Uid = account.Uid
	}

	if account.Uuid != "" {
		a.Uuid = account.Uuid
	}

	if account.Nickname != "" {
		a.Nickname = account.Nickname
	}

	if account.Profession != "" {
		a.Profession = account.Profession
	}

	if account.RG != "" {
		a.RG = account.RG
	}

	if account.Birthday != "" {
		a.Birthday = account.Birthday
	}

	if fmt.Sprintf(" %T", account.Options) == "bool" {
		a.Options = account.Options
	}

	if account.Gender != "" {
		a.Gender = account.Gender
	}

	if account.Status != Undefined {
		a.Status = account.Status
	}

	if account.CreatedAt != "" {
		a.CreatedAt = account.CreatedAt
	}

	if account.UpdatedAt != "" {
		a.UpdatedAt = account.UpdatedAt
	}

	return a
}
