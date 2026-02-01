package main

import (
	"battle-golang/internal"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Post("/api/calc", internal.Handler)
	log.Fatal(app.Listen(":8080"))
}
