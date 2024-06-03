package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}

const (
	TOKEN_KEY = "music_platform"
	ISSUER    = "music_platform"
)

// 签发用户token
func GenerateToken(id uint64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    ISSUER,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(TOKEN_KEY)
	return token, err
}

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return TOKEN_KEY, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
