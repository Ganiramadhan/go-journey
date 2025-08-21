package migrations

import (
	"log"

	"go-journey/src/database"
	"go-journey/src/model"
)

func Migrate() {
	log.Println("🚀 Running migration...")

	err := database.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("❌ Migration failed: ", err)
	}

	log.Println("✅ Migration completed: User table created")
}
