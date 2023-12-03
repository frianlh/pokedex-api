package configs

import "github.com/gofiber/fiber/v2/middleware/cors"

// CorsConfig is
func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST, GET, HEAD, PUT, DELETE, PATCH, OPTIONS",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With",
		AllowCredentials: true,
	}
}
