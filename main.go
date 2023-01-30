package main

import (
	"context"
	"github.com/dev-hyunsang/golang-jwt-redis/database"
	"log"

	"github.com/dev-hyunsang/golang-jwt-redis/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)

	client, err := database.ConnectionMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[DONE] Create DataBase Table...!")
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalln(err)
	}

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}
}
