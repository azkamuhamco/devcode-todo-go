package activity

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	go app.Get("/activity-groups", GetActivities)
	go app.Post("/activity-groups", CreateActivity)
	go app.Get("/activity-groups/:id", GetActivity)
	go app.Patch("/activity-groups/:id", UpdateActivity)
	go app.Delete("/activity-groups/:id", DeleteActivity)
}
