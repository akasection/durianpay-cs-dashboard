package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtSecret []byte

type Claims struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

func GenerateToken(username string, roles *[]string) (string, *Claims, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		Roles:    *roles,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    "gin-cs-dashboard",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, tokenErr := token.SignedString(JwtSecret)
	return signed, &claims, tokenErr
}

func ParseToken(tokenString string) (*Claims, error) {
	// TODO: implement
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GetTokenFromRequest(c *gin.Context) (string, error) {
	token := c.Query("token")
	headerToken := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	if len(headerToken) > len(prefix) && headerToken[:len(prefix)] == prefix {
		token = headerToken[len(prefix):]
	}

	if token == "" {
		return "", errors.New("authorization token not found")
	}

	return token, nil
}
