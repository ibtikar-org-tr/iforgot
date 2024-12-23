package services

import (
	"time"

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

func CheckLastSent(session *models.Session) bool {
	if session.Frequency == 0 {
		return true
	}
	if time.Since(session.LastSent) > 24*time.Hour {
		session.Frequency = 0
		repositories.UpdateSession(session)
		return true
	}

	waitTimes := []int{1, 5, 15, 30, 60} // in minutes
	waitTime := waitTimes[session.Frequency-1]
	if session.Frequency > len(waitTimes) {
		waitTime = waitTimes[len(waitTimes)-1]
	}

	lastSent := session.LastSent

	if lastSent.Add(time.Duration(waitTime) * time.Minute).Before(time.Now()) {
		repositories.UpdateSession(session)
		return true
	}

	return false
}
