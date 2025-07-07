package repositories

import (
	"gorm.io/gorm"
)

var UserRepository StructUserRepository
var AuthTokenRepository StructAuthTokenRepository

func InitRepositories(db *gorm.DB) {
	UserRepository = StructUserRepository{db: db}
	AuthTokenRepository = StructAuthTokenRepository{db: db}
}
