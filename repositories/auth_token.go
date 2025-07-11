package repositories

import (
	"gorm.io/gorm"

	"github.com/natanfds/vtt_odisseia/models"
)

type StructAuthTokenRepository struct {
	db *gorm.DB
}

func (a *StructAuthTokenRepository) CreateToken(user models.User, generatedToken string) error {

	token := models.AuthToken{
		Token:  generatedToken,
		UserID: user.ID,
	}

	result := a.db.Create(&token)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *StructAuthTokenRepository) GetTokenByID(userID string) (string, error) {
	var token models.AuthToken

	result := a.db.Where("user_id = ?", userID).First(&token)

	if result.Error != nil {
		return "", result.Error
	}

	return token.Token, nil
}

func (a *StructAuthTokenRepository) UpdateToken(newToken string, user models.User) error {
	result := a.db.Model(&models.AuthToken{}).Where("user_id = ?", user.ID).Update("token", newToken)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *StructAuthTokenRepository) DeleteToken(userId string) error {
	result := a.db.Delete(&models.AuthToken{}, "user_id = ?", userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
