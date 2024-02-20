package models

import (
	"time"

	"gorm.io/gorm"
)

type Rental struct {
	gorm.Model
	UserId     int `json:"user_id"`
	User       User
	BookId     int `json:"book_id"`
	Book       Book
	RentalDate time.Time `json:"rental_date"`
	ReturnData time.Time `json:"return_date"`
}
