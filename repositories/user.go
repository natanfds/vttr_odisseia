package repositories

import (
	"gorm.io/gorm"

	"github.com/natanfds/vtt_odisseia/dtos"
	"github.com/natanfds/vtt_odisseia/models"
	"github.com/natanfds/vtt_odisseia/utils"
)

type StructUserRepository struct {
	db *gorm.DB
}

func (u *StructUserRepository) CreateUser(data dtos.CreateUser) error {
	hash, err := utils.CreateHash(data.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Username:    data.Username,
		DisplayName: data.DisplayName,
		Hash:        hash,
		Email:       data.Email,
	}
	result := u.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *StructUserRepository) GetUser(data dtos.GetUser) (models.User, error) {
	var users models.User

	searchArgs := &models.User{}
	if data.DisplayName != "" {
		searchArgs.DisplayName = data.DisplayName
	}
	if data.Email != "" {
		searchArgs.Email = data.Email
	}
	if data.Username != "" {
		searchArgs.Username = data.Username
	}
	result := u.db.Where(searchArgs).First(&users)

	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (u *StructUserRepository) UpdateUser(data dtos.UpdateUser, id int) error {
	result := u.db.Model(&models.User{}).Where("id = ?", id).Updates(data)
	return result.Error
}

func (u *StructUserRepository) DeleteUser(id int) error {
	result := u.db.Delete(&models.User{}, id)
	return result.Error
}
