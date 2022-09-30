package db

import (
	"log"
	"os"

	"github.com/rudikurniawan99/go-api-4/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	connection := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		log.Println("===")
		log.Println("can't connect to db", err)
		log.Println("===")
	} else {
		log.Println("===")
		log.Println("success to connect to db")
		log.Println("===")
	}
	db.AutoMigrate(&model.User{})

	return db
}
