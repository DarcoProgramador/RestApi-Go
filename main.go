package main

import (
	"gorm/handlers"
	"gorm/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//Migrando Modelo
	models.MigrarUser()

	//Rutas
	app := fiber.New()

	//EndPoint
	app.Use(logger.New())

	app.Get("/api/user", handlers.GetUsers)
	app.Get("/api/user/:id<int>", handlers.GetUser)
	app.Post("/api/user", handlers.CreateUser)
	app.Put("/api/user/:id<int>", handlers.UpdateUser)
	app.Delete("/api/user/:id<int>", handlers.DeleteUser)

	log.Fatal(app.Listen(":3000"))

}
