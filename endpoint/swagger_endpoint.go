package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetSwaggerEndpoint(fiberApp *fiber.App) {
	fiberApp.Get("/swagger-ui/*", swagger.New(swagger.Config{
		URL:          "/swagger-ui/doc.json",
		DocExpansion: "none",
	}))
}
