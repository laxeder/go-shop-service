package user

import "github.com/laxeder/go-shop-service/pkg/modules/str"

type UserDocument struct {
	Uuid        string `json:"uuid,omitempty"`
	OldDocument string `json:"old_document,omitempty"`
	NewDocument string `json:"new_document,omitempty"`
}

func (u *UserDocument) GenerateDocument() {
	u.OldDocument = str.DocumentPad(u.OldDocument)
	u.NewDocument = str.DocumentPad(u.NewDocument)
}
