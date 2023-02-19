package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt私钥
var jwtSecret = []byte("jwt_secret")

type Claims struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken 签发一个token
func GenerateToken(id uint, userName string) (string, error) {
	now := time.Now()
	expire := now.Add(7 * 24 * time.Hour)
	claims := Claims{
		ID:       id,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    "todo_list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err == nil {
		return token, nil
	}
	return "", err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	parseClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if parseClaims != nil && err == nil {
		if claims, ok := parseClaims.Claims.(*Claims); ok && parseClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
