package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConf() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout:                  time.Second * time.Duration(readTimeoutSecondsCount),
		DisableStartupMessage:        true,
		DisableKeepalive:             true,
		StrictRouting:                true,
		DisablePreParseMultipartForm: true,
		JSONEncoder:                  json.Marshal,
		JSONDecoder:                  json.Unmarshal,
	}
}
