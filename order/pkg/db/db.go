package db

import (
	"log"

	"github.com/Brenopms/grpc-learning/order/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Initializing Migrations")

	db.AutoMigrate(&models.Order{})

	log.Println("Migrations are done")

	return Handler{db}
}
