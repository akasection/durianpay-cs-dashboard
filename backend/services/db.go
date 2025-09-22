package services

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbPath = "db/local.sqlite"
var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DbPath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	} else {
		// Auto migrate

		DB = db
		log.Println("Loaded db from", DbPath)
	}

	return db, err
}
