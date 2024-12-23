package main

import (
	"github.com/ibtikar-org-tr/iforgot/initializers"
	"github.com/ibtikar-org-tr/iforgot/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Session{})
	initializers.DB.AutoMigrate(&models.ID{})
}
