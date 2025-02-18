package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

const SWAGGER = "/swagger-ui"

func SetSwaggerEndpoint(fiberApp *fiber.App) {
	fiberApp.Get(SWAGGER+"/*", swagger.New(swagger.Config{
		URL:          SWAGGER + "/doc.json",
		DocExpansion: "none",
	}))
}
