package services

import (
	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/models"
	"github.com/ibtikar-org-tr/iforgot/repositories"
)

func GetSessionByIP(c *gin.Context) (*models.Session, error) {
	ip := c.ClientIP()
	session, err := repositories.GetSessionByIP(ip)
	if err != nil {
		return nil, err
	}
	return session, nil
}
