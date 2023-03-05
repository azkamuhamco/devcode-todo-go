package activity

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/activity-groups", GetActivities)
	app.Post("/activity-groups", CreateActivity)
	app.Get("/activity-groups/:id", GetActivity)
	app.Patch("/activity-groups/:id", UpdateActivity)
	app.Delete("/activity-groups/:id", DeleteActivity)
}
