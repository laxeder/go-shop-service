package jwt

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
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

func (c *Claims) InjectMap(claimsMap any) *Claims {
	var log = logger.New()

	b, err := json.Marshal(claimsMap)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao transformar claims map en byte %s", claimsMap)
		return c
	}

	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao transformar byte em claims map: %s", claimsMap)
		return c
	}

	return c
}
