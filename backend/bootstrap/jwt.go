package bootstrap

import (
	"os"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
)

func SetupJWT() {
	secret := os.Getenv("JWT_SECRET")
	util.JwtSecret = []byte(secret)
}
