package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Description   string    `json:"description"`
	Publication   time.Time `json:"publication"`
	NumberOfPages int       `json:"number_of_pages"`
}
