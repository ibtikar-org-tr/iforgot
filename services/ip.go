package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibtikar-org-tr/iforgot/models"
	"github.com/ibtikar-org-tr/iforgot/repositories"
)

func GetSession(c *gin.Context) (*models.Session, error) {
	ip := c.ClientIP()
	session, err := repositories.GetSessionByIP(ip)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func GetSessionByIP(ip string) (*models.Session, error) {
	session, err := repositories.GetSessionByIP(ip)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func StoreSession(ip string) error {
	session, err := repositories.GetSessionByIP(ip)
	if err != nil {
		// If session does not exist, create a new one
		newSession := &models.Session{
			IP:        ip,
			LastSent:  time.Now(),
			Frequency: 0,
		}
		return repositories.CreateSession(newSession)
	}

	// If session exists, update the LastSent to now
	session.LastSent = time.Now()
	return repositories.UpdateSession(session)
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
