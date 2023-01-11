package tokens

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewUUID(args ...string) string {
	nonce := time.Now().UnixMilli()

	var token string
	for _, arg := range args {
		token = fmt.Sprintf("%v%v", token, arg)
	}

	uuidModel, _ := uuid.Parse("b9cfdb9d-f741-4e1f-89ae-fac6b2a5d740")
	hashByte32 := sha256.Sum256([]byte(fmt.Sprintf("%v%v", nonce, token)))
	hashBase := b64.StdEncoding.EncodeToString(hashByte32[:])

	return uuid.NewSHA1(uuidModel, []byte(hashBase)).String()
}
