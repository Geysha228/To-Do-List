package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(user, password, host, port, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345678 dbname=tododb port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return db, nil
}
