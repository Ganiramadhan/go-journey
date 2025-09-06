package router

import (
	_ "go-journey/src/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func DocsRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
}
