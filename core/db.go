package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=tasks port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	errAM := db.AutoMigrate(&Task{})
	if errAM != nil {
		return nil, errAM
	}

	return db, nil
}
