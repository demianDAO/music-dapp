package utils

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	TOKEN_KEY = []byte("music_platform")
	ISSUER    = "music_platform"
)

type CustomClaims struct {
	Address string `json:"address"`
	*jwt.StandardClaims
}

// 签发用户token
func GenerateToken(address string) (string, error) {

	claims := CustomClaims{
		Address: address,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    ISSUER,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(TOKEN_KEY)
	if err != nil {
		log.Print("token generate error:", err)
	}
	return token, err
}

// 验证用户token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return TOKEN_KEY, nil
	})

	if err != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
