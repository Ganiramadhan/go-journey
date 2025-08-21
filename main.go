package main

import (
	"log"

	"github.com/ganiramadhan/go-fiber-app/internal/handler"
	"github.com/ganiramadhan/go-fiber-app/internal/repository"
	"github.com/ganiramadhan/go-fiber-app/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// === User ===
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app.Get("/users", userHandler.GetUsers)
	app.Post("/users", userHandler.CreateUser)

	// Start server with log + error handling
	log.Println("🚀 Server running at http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
