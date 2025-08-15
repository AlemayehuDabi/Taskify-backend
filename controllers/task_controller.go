package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AlemayehuDabi/Taskify-backend/db"
	"github.com/AlemayehuDabi/Taskify-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var taskCollection = db.Client.Database("taskify").Collection("task")

// get task
func GetTasks(c *fiber.Ctx) error {

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	cursor, err := taskCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var tasks []models.Task

	if err := cursor.All(ctx, &tasks); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"msg": "successfully fetch todos", "todos": tasks})
}

// create task
 func CreateTask(c *fiber.Ctx) error {
		
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancle()

	cursor, err := taskCollection.Find(ctx, bson.M{})

	
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
	}

// toggle complete
func ToggleComplete(c *fiber.Ctx) error  {
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
		
	}

// update task
func UpdateTask(c *fiber.Ctx) error {
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

	}

	// delete task
func DeleteTask(c *fiber.Ctx) error  {
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
		
	}