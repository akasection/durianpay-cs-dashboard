package bootstrap

import (
	"os"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
)

func SetupJWT() {
	util.JwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
