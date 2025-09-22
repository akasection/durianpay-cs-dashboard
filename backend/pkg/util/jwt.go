package util

import (
	// "time"

	"github.com/dgrijalva/jwt-go"
)

// var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// func GenerateToken(username, password, role string) (string, error) {
// 	nowTime := time.Now()

// }

// func ParseToken(tokenString string) (*Claims, error) {

// }
