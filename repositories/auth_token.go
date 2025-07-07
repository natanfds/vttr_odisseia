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

func (a *StructAuthTokenRepository) GetToken() {}

func (a *StructAuthTokenRepository) UpdateToken() {}

func (a *StructAuthTokenRepository) DeleteToken() {}
