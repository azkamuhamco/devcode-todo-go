package activity

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	api := app.Group("/activity-groups")
	api.Get("", GetActivities)
	api.Post("", CreateActivity)
	api.Get("/:id", GetActivity)
	api.Patch("/:id", UpdateActivity)
	api.Delete("/:id", DeleteActivity)
}
