package repositories

import (
	"github.com/ibtikar-org-tr/iforgot/initializers"
	"github.com/ibtikar-org-tr/iforgot/models"
)

func GetSessionByIP(ip string) (*models.Session, error) {
	session := &models.Session{}
	if err := initializers.DB.Where("ip = ?", ip).First(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func CreateSession(session *models.Session) error {
	if err := initializers.DB.Create(session).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSession(session *models.Session) error {
	if err := initializers.DB.Save(session).Error; err != nil {
		return err
	}
	return nil
}
