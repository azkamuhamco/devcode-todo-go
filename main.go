package main

import (
	"github.com/gofiber/fiber/v2"

	"devcode-todo-go/internal/apps/activity"
	"devcode-todo-go/internal/apps/todo"
	"devcode-todo-go/internal/models"
	"devcode-todo-go/pkg/configs"
	"devcode-todo-go/pkg/database"
	"devcode-todo-go/pkg/utils"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	// Open DB connection
	url := utils.Url()
	gormC := configs.GormConfig()
	database.DBConn, _ = gorm.Open(mysql.Open(url), gormC)

	// Auto Migrate
	database.DBConn.AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)
}

func main() {
	// Variable
	konpig := configs.FiberConf()
	app := fiber.New(konpig)

	// Route
	activity.Route(app)
	todo.Route(app)

	// Listen Port
	app.Listen(":3030")
}
