package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	IP        string    `json:"ip"`
	LastCheck time.Time `json:"last_check"`
}
