package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/services"
)

func SearchHandler(c *gin.Context) {
	value := c.Query("value")
	typeOfValue := c.Query("type")

	ip := c.ClientIP()

	// log the request
	fmt.Printf("Searching for value: %s, type: %s\n", value, typeOfValue)

	result, err := services.SearchMain(value, typeOfValue, ip)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": result})
}
