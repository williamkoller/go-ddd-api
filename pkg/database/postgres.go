package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var db *gorm.DB
	var err error

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Successfully connected to the database!")
			return db, nil
		}

		log.Printf("🔄 Failed to connect to database. Retrying (%d/%d)...", i+1, maxRetries)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("❌ Could not connect to the database after %d retries: %v", maxRetries, err)
	return nil, nil
}
