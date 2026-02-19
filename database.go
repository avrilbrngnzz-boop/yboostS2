package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {

		dsn = "user=postgres password=postgres host=localhost port=5432 dbname=yboost_test sslmode=disable"

	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erreur de connexion à la base: %w", err)
	}

	return db, nil
}
