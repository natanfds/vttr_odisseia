package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex:idx_user_name"`
	DisplayName string `gorm:"uniqueIndex:idx_display_name"`
	Hash        string
	Email       string `gorm:"uniqueIndex:idx_user_email"`
}
