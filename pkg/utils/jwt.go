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
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
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
	log.Print("tokenString:", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return TOKEN_KEY, nil
	})

	if err != nil {
		log.Print("token parse error:", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
