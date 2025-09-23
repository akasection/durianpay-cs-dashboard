package main

import (
	"os"

	"github.com/akasection/durianpay-cs-dashboard/backend/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GO_ENV") != "production" || os.Getenv("GIN_MODE") != "release" {
		godotenv.Load()
	}

	bootstrap.Initialize()
}
