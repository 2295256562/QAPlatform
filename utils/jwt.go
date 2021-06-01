package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JWTSECRET = "23347$040412"
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func GenerateToken(id int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * time.Hour)
	claims := jwt.MapClaims{
		"id":   id,
		"name": username,
		"exp":  expireTime, // 过期时间，必须设置,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(JWTSECRET))
	return token, err
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
