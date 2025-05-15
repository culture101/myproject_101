// cmd/main.go

package main

import (
	"github.com/culture101/myproject_101.git/handlers"
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2" // 1. импортируем пакет шаблонов
)

func main() {

	engine := html.New("./views", ".html") //Создаю новый движок

	database.ConnectDb()

	app := fiber.New(fiber.Config{
		Views:       engine,         // new config
		ViewsLayout: "layouts/main", // add this to config
	})

	setupRoutes(app)
	app.Static("/", "./public")

	// Set up 404 page
	app.Use(handlers.NotFound)

	app.Listen(":3000")
}
