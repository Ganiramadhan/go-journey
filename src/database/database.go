package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	customLogger := &CustomLogger{LogLevel: logger.Info}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}

	DB = db
	log.Println("✅ Connected to Database")
}
