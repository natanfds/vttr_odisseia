package services

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/natanfds/vtt_odisseia/models"
)

func StartDatabase() (*gorm.DB, error) {
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.AuthToken{},
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, model := range modelsToMigrate {
		db.AutoMigrate(model)
	}

	return db, nil
}
