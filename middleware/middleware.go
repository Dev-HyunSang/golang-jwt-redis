package middleware

import (
	"github.com/dev-hyunsang/golang-jwt-redis/cmd"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/join", cmd.JoinUserHandler)
	auth.Post("/login", cmd.LoginUserHandler)

	todo := app.Group("/todo")
	todo.Post("/create", cmd.CreateToDo)
}
