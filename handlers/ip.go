package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/models"
	"github.com/ibtikar-org-tr/iforgot/repositories"
)

func GetIP(c *gin.Context) *models.Session {
	ip := c.ClientIP()
	session, err := repositories.GetSessionByIP(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	return session
}
