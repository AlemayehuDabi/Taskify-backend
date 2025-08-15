package routes

import (
	"github.com/AlemayehuDabi/Taskify-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func TaskRoute(app *fiber.App) {

	api := app.Group("/api/todos")

	// get to do route
	api.Get("/", controllers.GetTasks)

	// create a todo
	api.Post("/", controllers.CreateTask)

	// toggle completed
	api.Patch("/:id", controllers.ToggleComplete)

	// update a todo
	api.Put("/:id", controllers.UpdateTask)

	// delete a todo
	api.Delete("/:id", controllers.DeleteTask)

}
