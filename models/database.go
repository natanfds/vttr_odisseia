package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func (d *database) StartDatabase() error {
	modelsToMigrate := []interface{}{
		&User{},
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	for _, model := range modelsToMigrate {
		db.AutoMigrate(model)
	}

	d.db = db

	return nil
}

var Database = &database{}
