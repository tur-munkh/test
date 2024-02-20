// @/main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()
	var handlers UserHandler
	app.Get("/users", handlers.Getusers)
	app.Get("/users/:id", handlers.Getuser)
	app.Post("/users", handlers.Adduser)
	app.Put("/users/:id", handlers.Updateuser)
	app.Delete("/users/:id", handlers.Removeuser)

	log.Fatal(app.Listen(":5432"))
}
