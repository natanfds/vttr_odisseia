package repositories

import (
	"gorm.io/gorm"
)

var UserRepository StructUserRepository

func InitRepositories(db *gorm.DB) {
	UserRepository = StructUserRepository{db: db}
}
