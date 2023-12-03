package configs

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"time"
)

// LoggerConfig is
func LoggerConfig() logger.Config {
	return logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stderr,
	}
}
