package db

import (
	"api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Current.DatabaseDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
