package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JWTSECRET           = "23347$040412"
	TokenExpireDuration = time.Hour * 2
)

type MyClaims struct {
	UserId   int
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id int, username string) (string, error) {
	claims := MyClaims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "platform",                                 // 签发人
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(JWTSECRET))
	return token, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(JWTSECRET), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
