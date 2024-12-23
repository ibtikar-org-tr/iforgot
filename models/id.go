package models

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	gorm.Model
	Number    int       `json:"number"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	LastSent  time.Time `json:"last_check"`
	Frequency int       `json:"frequency"`
}
