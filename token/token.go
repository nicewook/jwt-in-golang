package token

import (
	"fmt"
	"jwt-in-golang/model"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	privateKeyForJWT = "SecretTokenSecretToken"
	ip               = "192.168.0.203"
)

func GenerateToken(claims *model.JWTClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(privateKeyForJWT))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getTokenFromString(tokenString string, claims *model.JWTClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Remember to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected sigining method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(privateKeyForJWT), nil
	})
}

func VerifyToken(tokenString, origin string) (bool, *model.JWTClaims) {
	claims := &model.JWTClaims{}
	token, _ := getTokenFromString(tokenString, claims)
	if token.Valid {
		if err := claims.Valid(); err == nil {
			return true, claims
		}
	}
	return false, claims
}

func GetClaims(tokenString string) model.JWTClaims {
	claims := &model.JWTClaims{}

	_, err := getTokenFromString(tokenString, claims)
	if err != nil {
		log.Println(err)
		return *claims
	}
	return *claims

}
