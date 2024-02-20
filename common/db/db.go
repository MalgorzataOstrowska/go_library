package db

import (
	"log"

	"library/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	// dsn := "host=localhost user=postgres password=root dbname=library port=5432 sslmode=disable TimeZone=CET"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{}, &models.User{}, &models.Rental{})

	return db
}
