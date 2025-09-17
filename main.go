package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-journey/src/database"
	"go-journey/src/database/migrations"
	"go-journey/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"go-journey/src/docs"
)

// @title           User API
// @version         1.0
// @description     API service for managing users
// @host            127.0.0.1:8080
// @BasePath        /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer {your token}" (without quotes)
func main() {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		if err := godotenv.Load(".env.production"); err != nil {
			log.Println("⚠️ No .env.production file found, using system env")
		} else {
			log.Println("✅ Loaded .env.production")
		}
	} else {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("⚠️ No .env file found, using system env")
		} else {
			log.Println("✅ Loaded .env")
		}
	}

	// Database connection
	database.ConnectDB()
	migrations.Migrate()

	// Fiber app config
	app := fiber.New(fiber.Config{
		AppName:       "User API v1.0",
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(recover.New())

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Routes
	router.UserRoutes(app)
	router.AuthRoutes(app)
	router.DocsRoutes(app)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}
	addr := fmt.Sprintf(":%s", port)

	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = fmt.Sprintf("127.0.0.1:%s", port)
	}
	docs.SwaggerInfo.Host = swaggerHost

	log.Printf("📘 Swagger running at http://%s/docs/index.html", swaggerHost)

	// Start server
	go func() {
		log.Printf("🚀 Server running at 0.0.0.0%s", addr)
		if err := app.Listen(addr); err != nil {
			log.Fatalf("❌ Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited properly")
}
