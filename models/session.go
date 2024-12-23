package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	IP        string    `json:"ip"`
	LastSent  time.Time `json:"last_sent"`
	Frequency int       `json:"frequency"`
}
