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
	LastCheck time.Time `json:"last_check"`
}
