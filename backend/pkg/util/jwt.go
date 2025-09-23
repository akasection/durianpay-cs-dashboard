package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret []byte

type Claims struct {
	Username  string   `json:"username"`
	HPassword string   `json:"password"`
	Roles     []string `json:"roles"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username:  username,
		HPassword: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	// TODO: implement
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
