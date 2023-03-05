package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	"devcode-todo-go/internal/apps/activity"
	"devcode-todo-go/internal/apps/todo"
	"devcode-todo-go/internal/models"
	"devcode-todo-go/pkg/configs"
	"devcode-todo-go/pkg/database"
	"devcode-todo-go/pkg/middleware"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	// Variabel to connect DB
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"))

	// Open DB connection
	database.DBConn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Auto Migrate
	database.DBConn.AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)
}

func setupRoutes(app *fiber.App) {
	activity.Route(app)
	todo.Route(app)
}

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Route
	setupRoutes(app)

	// Listen Port
	app.Listen(":3030")
}
