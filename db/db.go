package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the database
	db.AutoMigrate( /* add the models */ )

	return db, nil
}
