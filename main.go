package main

import (
	"fmt"
	"log"
	"os"

	"go-journey/src/database"
	"go-journey/src/database/migrations"
	"go-journey/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "go-journey/src/docs"
)

// @title           User API
// @version         1.0
// @description     API service for managing users
// @host            localhost:8080
// @BasePath        /
func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env")
	}

	// Connect DB
	database.ConnectDB()

	// Run migrations
	migrations.Migrate()

	// Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Routes
	router.UserRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Get PORT from .env or default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("🚀 Server running at http://localhost%s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
