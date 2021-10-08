package model

import (
	"errors"
	"time"

	"crypto/subtle"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	CompanyID string
	Username  string
	Roles     []int
	jwt.StandardClaims
}

const ip = "192.168.0.203"

func (claims JWTClaims) Valid() error {
	nowUnix := time.Now().Unix() // https://play.golang.org/p/22u_Ecjs3eJ
	if claims.VerifyExpiresAt(nowUnix, true) && claims.VerifyIssuer(ip, true) {
		return nil
	}
	return errors.New("Token is invalid")
}

func (claims JWTClaims) VerifyAudience(origin string) bool {
	return subtle.ConstantTimeCompare([]byte(claims.Audience), []byte(origin)) == 1
}
