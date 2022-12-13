package jwt

import (
	"time"

	"github.com/google/uuid"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
)

const (
	ISS string = "57408a9a-f382-4c96-9d52-54aecfdf9f3f"
	AUD string = "57408a9a-f382-4c96-9d52-54aecfdf9f3f"
)

type Claims struct {
	Jti  string `json:"jti,omitempty"`  // token ID
	Iss  string `json:"iss,omitempty"`  // emissor ID
	Sub  string `json:"sub,omitempty"`  // user ID
	Aud  string `json:"aud,omitempty"`  // audencia ID
	Iat  string `json:"iat,omitempty"`  // data da emissao
	Nbf  string `json:"nbf,omitempty"`  // data a partir da qual o token será aceito
	Exp  string `json:"exp,omitempty"`  // data da expiracao
	Data any    `json:"data,omitempty"` // dado adicional
}

func (c *Claims) autoload(session time.Duration) *Claims {

	jti := uuid.New().String()

	c.Jti = jti
	c.Iss = ISS
	c.Sub = ""
	c.Aud = AUD
	c.Iat = date.NowUTC()
	c.Nbf = date.NowUTC()
	c.Exp = date.NowUTCAddMinutes(session)

	return c

}
