package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {

	err := godotenv.Load()

	if(err != nil){
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func (c *fiber.Ctx) error {
		return 	c.Status(200).JSON(fiber.Map{"msg" : "successfully fetch todos", "todos" : todos})
		})

	// create a todo
	app.Post("/api/todos", func (c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if(todo == &Todo{}){
			return c.Status(401).JSON(fiber.Map{"error": "Todo is required"})
		}

		if (todo.Body == "") {
			return c.Status(401).JSON(fiber.Map{"error": "Todo body is required"})
		}

		// increase id value
		todo.ID = len(todos) + 1

		// put the data in the todos array
		todos = append(todos, *todo)

		return c.Status(200).JSON(fiber.Map{"msg": "successfully", "todos": todos})
	})

	// mark as completed
	app.Patch("/api/todos/:id", func (c *fiber.Ctx) error  {
		id := c.Params("id")

		if(id == ""){
			return c.Status(401).JSON(fiber.Map{"error": "Todo ID is required in order to mark as completed",})
		}

		for i:= range todos {
			if(fmt.Sprint(todos[i].ID) == id){
				todos[i].Completed = true
				return c.Status(200).JSON(fiber.Map{"msg": "successfully mark as completed", "todos": todos})
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error" : "unable to mark as completed"})
		
	})


	// update a todo
	app.Put("/api/todos/:id", func(c *fiber.Ctx) error {
		id :=  c.Params("id")

		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if(todo == &Todo{}){
			return c.Status(401).JSON(fiber.Map{"error": "Todo is required inorder to update"})
		}

		if(id == ""){
			return c.Status(401).JSON(fiber.Map{"error": "Todo ID is required in order to update",})
		}

		if(todo.Body == ""){
			return c.Status(401).JSON(fiber.Map{"error": "Todo Body is required in order to update",})
		}

		for i := range todos {
			if (fmt.Sprint(todos[i].ID) == id){
				todos[i].Completed = (*todo).Completed
				todos[i].Body = (*todo).Body
				return c.Status(200).JSON(fiber.Map{"msg": "successfully update a todo", "todos": todos})
			}
		} 

		return c.Status(404).JSON(fiber.Map{
			"error" : "unable to update"})

	})


	// delete a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error  {
		id := c.Params("id")

			if(id == ""){
			return c.Status(401).JSON(fiber.Map{"error": "Todo ID is required in order to update",})
		}

		for i := range todos {
			if(fmt.Sprint(todos[i]) == id){
				todos = append(todos[: i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "successfully delete the todo", "todos": todos})
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error" : "unable to delete a todo"})
		
	})

	log.Printf("Server running on port %s\n", port)
	log.Fatal(app.Listen(":"+port))
}