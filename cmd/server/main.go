package main

import (
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

	

	log.Printf("Server running on port %s\n", port)
	log.Fatal(app.Listen(":"+port))
}