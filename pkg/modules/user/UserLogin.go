package user

type UserLogin struct {
	Email    string `json:"email,omitempty" redis:"email,omitempty"`
	Password string `json:"password,omitempty" redis:"password,omitempty"`
	Salt     string `json:"salt,omitempty" redis:"salt,omitempty"`
}
