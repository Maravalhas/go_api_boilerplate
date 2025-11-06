package db

import (
	"api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DatabaseDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
