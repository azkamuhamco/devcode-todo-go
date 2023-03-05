package todo

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	go app.Get("/todo-items", GetTodos)
	app.Post("/todo-items", CreateTodo)
	go app.Get("/todo-items/:id", GetTodo)
	app.Patch("/todo-items/:id", UpdateTodo)
	go app.Delete("/todo-items/:id", DeleteTodo)
}
