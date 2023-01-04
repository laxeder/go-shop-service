package user

type UserPassword struct {
	Uuid            string `json:"uuid,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	OldPassword     string `json:"old_password,omitempty"`
}
