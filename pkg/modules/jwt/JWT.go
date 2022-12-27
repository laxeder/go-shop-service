package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
	"github.com/laxeder/go-shop-service/pkg/utils"
)

type JWT struct {
	Token        string `json:"token,omitempty"`
	Claims       Claims `json:"claims,omitempty"`
	Validated    bool   `json:"validated,omitempty"`
	Signed       bool   `json:"signed,omitempty"`
	Session      int    `json:"session,,omitempty"`
	PublicKey    string `json:"public_key,omitempty"`
	PrivateKey   string `json:"private_key,omitempty"`
	SignatureKey string `json:"signature_key,omitempty"`
}

func New(data any) (jwtData *JWT) {

	var log = logger.New()

	jwtData = &JWT{}

	// inicia o claims
	claims := Claims{}
	claims.autoload(120)

	// Create the Claims
	claimsJwt := jwt.MapClaims{
		"jti":  claims.Jti,
		"iss":  claims.Iss,
		"sub":  claims.Sub,
		"aud":  claims.Aud,
		"iat":  date.UTCToTime(claims.Iat).Unix(),
		"nbf":  date.UTCToTime(claims.Nbf).Unix(),
		"exp":  date.UTCToTime(claims.Exp).Unix(),
		"data": data,
	}

	// cria um objeto para o token
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsJwt)

	jwtData.Claims = claims
	jwtData.Credentials()

	// assina um token
	// token, err := tokenData.SignedString([]byte(jwtData.PrivateKey))
	token, err := tokenData.SignedString([]byte("secret"))
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possivel assinar o token. %v", err)
	}

	jwtData.Token = token

	return
}

func (j *JWT) Result() fiber.Map {
	return fiber.Map{"token": j.Token}
}

func (j *JWT) Credentials() *JWT {

	read := utils.ReadFile{}

	j.PublicKey = read.Run("./keys/publickey.pem")
	j.PrivateKey = read.Run("./keys/privatekey.pem")
	j.SignatureKey = str.MixStrings(j.PrivateKey, j.PublicKey)

	return j
}

func Signature() string {

	read := utils.ReadFile{}

	PublicKey := read.Run("./keys/publickey.pem")
	// PrivateKey := read.Run("./keys/privatekey.pem")
	// return str.MixStrings(PrivateKey, PublicKey)
	return PublicKey
}

func Info(token string) *JWT {
	return &JWT{Token: token}
}

//  JWT = jwt.New(user) -> criar um token
// token = JWT.Token
//  JWT = 	jwt.Info(token) -> get uma struct
//  claims = jwt.Get(token) -> pegar o claims no token
//	valid = jtw.Valid(token) -> validacao/expirado
//  token = jwt.Reresh(token, ?:clamis) -> atualiza um token e o claims (o claims é opcional)
