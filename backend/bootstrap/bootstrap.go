package bootstrap

import (
	"os"

	"github.com/akasection/durianpay-cs-dashboard/backend/routers"
	"github.com/akasection/durianpay-cs-dashboard/backend/services"
)

func Initialize() {
	SetupJWT()
	db, err := services.ConnectDB()
	if err != nil {
		panic("failed to initialize database")
	}
	MigrateDB(db)
	router := routers.SetupRouter()

	// TODO: setup redis?

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8081"
	}
	router.Run(listenAddr)
}
