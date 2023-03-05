package todo

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/todo-items", GetTodos)
	app.Post("/todo-items", CreateTodo)
	app.Get("/todo-items/:id", GetTodo)
	app.Patch("/todo-items/:id", UpdateTodo)
	app.Delete("/todo-items/:id", DeleteTodo)
}
