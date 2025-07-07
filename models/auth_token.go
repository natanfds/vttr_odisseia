package models

import "gorm.io/gorm"

type AuthToken struct {
	gorm.Model
	Token  string
	UserID uint
	User   User `gorm:"foreignKey:UserID;references:ID"`
}
