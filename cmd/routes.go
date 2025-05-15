package main

import (
	"github.com/culture101/myproject_101.git/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// Обновленный маршрут: при GET-запросе на "/" вызывается функция ListFacts
	app.Get("/", handlers.ListFacts)
	app.Get("/fact", handlers.NewFactView) // Add new route for new view
	app.Post("/fact", handlers.CreateFact) // Регистрируем POST запрос по адресу "/fact"

	// Add new route to show single Fact, given `:id`
	app.Get("/fact/:id", handlers.ShowFact)

	// Display `Edit` form
	app.Get("/fact/:id/edit", handlers.EditFact)
	// Update fact
	app.Patch("/fact/:id", handlers.UpdateFact)

	app.Delete("/fact/:id", handlers.DeleteFact)

}
