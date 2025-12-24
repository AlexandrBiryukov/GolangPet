package db

import (
	"golang/internal/services"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost  user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Нет подключения к БД: %v", err)
	}

	if err := db.AutoMigrate(&services.RequestBody{}); err != nil {
		log.Fatalf("Мигрвция невозможна: %v", err)

	}
	return db, nil
}
