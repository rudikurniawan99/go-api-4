package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	connection := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(connection))

	if err != nil {
		log.Println("can't connect to db", err)
	}
	log.Println("===")
	log.Println("success to connect to db")
	log.Println("===")

	return db
}
