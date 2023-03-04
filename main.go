package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"devcode-todo-go/database"
	"devcode-todo-go/internal/apps/activity"
	"devcode-todo-go/internal/apps/todo"
	"devcode-todo-go/internal/models"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Variabel to connect DB
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"))

	// Open DB connection
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

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
	app := fiber.New(fiber.Config{
		DisableStartupMessage:        true,
		DisableKeepalive:             true,
		StrictRouting:                true,
		DisablePreParseMultipartForm: true,
	})
	initDB()
	setupRoutes(app)

	log.Fatal(app.Listen(":3030"))
}
