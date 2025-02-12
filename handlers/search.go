package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/services"
)

func SearchHandler(c *gin.Context) {
	typeOfValue := c.Query("type")
	var value string

	switch typeOfValue {
	case "number":
		value = c.Query("number")
	case "phone":
		value = c.Query("phone")
	case "mail":
		value = c.Query("mail")
	default:
		c.JSON(400, gin.H{"error": "Invalid type"})
		return
	}

	ip := c.ClientIP()

	// log the request
	fmt.Printf("Searching for value: %s, type: %s\n", value, typeOfValue)

	mail, err := services.SearchMain(value, typeOfValue, ip)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// hide part of the mail
	result := mail[:4] + "****" + mail[len(mail)-14:]

	// if the mail is shorter than 15 characters, show only 2 - 11
	if len(mail) < 15 {
		result = mail[:2] + "****" + mail[len(mail)-11:]
	}

	c.JSON(200, gin.H{
		"mail":    result,
		"message": "تمّ إرسال البيانات إلى البريد",
	})
}
