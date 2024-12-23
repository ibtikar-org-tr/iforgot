package repositories

import (
	"github.com/ibtikar-org-tr/iforgot/initializers"
	"github.com/ibtikar-org-tr/iforgot/models"
)

func CreateID(id *models.ID) error {
	if err := initializers.DB.Create(id).Error; err != nil {
		return err
	}
	return nil
}

func GetIDByEmail(email string) (*models.ID, error) {
	id := &models.ID{}
	if err := initializers.DB.Where("email = ?", email).First(id).Error; err != nil {
		return nil, err
	}
	return id, nil
}

func GetIDByPhone(phone string) (*models.ID, error) {
	id := &models.ID{}
	if err := initializers.DB.Where("phone = ?", phone).First(id).Error; err != nil {
		return nil, err
	}
	return id, nil
}

func GetIDByNumber(number string) (*models.ID, error) {
	id := &models.ID{}
	if err := initializers.DB.Where("number = ?", number).First(id).Error; err != nil {
		return nil, err
	}
	return id, nil
}

func UpdateID(id *models.ID) error {
	if err := initializers.DB.Save(id).Error; err != nil {
		return err
	}
	return nil
}
