package main

import (
	"context"
	"github.com/dev-hyunsang/golang-jwt-redis/auth"
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

	log.Println("[DONE] Create DataBase Table based Schema")
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalln(err)
	}

	redisClient := auth.RedisInit()
	result, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("[DONE] Redis PING %s", result)

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}
}
