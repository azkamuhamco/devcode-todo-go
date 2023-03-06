package todo

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	apt := app.Group("/todo-items")
	apt.Get("", GetTodos)
	apt.Post("", CreateTodo)
	apt.Get("/:id", GetTodo)
	apt.Patch("/:id", UpdateTodo)
	apt.Delete("/:id", DeleteTodo)
}
