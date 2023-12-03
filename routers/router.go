package routers

import (
	"github.com/frianlh/pokedex-api/configs"
	"github.com/frianlh/pokedex-api/routers/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoute is
func SetupRoute(config *configs.Config) *fiber.App {
	f := fiber.New()
	f.Use(cors.New(configs.CorsConfig()))
	f.Use(logger.New(configs.LoggerConfig()))

	// api route
	apiRoute := f.Group("/api")
	{
		v1 := apiRoute.Group("/v1")
		api.V1Route(v1, config)
	}

	return f
}
