package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartDatabase() (*gorm.DB, error) {
	modelsToMigrate := []interface{}{
		&User{},
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
