package main

import (
	"github.com/akasection/durianpay-cs-dashboard/backend/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bootstrap.Initialize()
}
