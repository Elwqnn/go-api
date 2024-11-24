package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	dsn := "postgres://axo:DTX70914!@localhost:5432/goapi?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	DB = db
}

// GetDB provides access to the initialized database
func GetDB() *gorm.DB {
	return DB
}
