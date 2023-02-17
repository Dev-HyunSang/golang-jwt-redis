package middleware

import (
	"github.com/dev-hyunsang/golang-jwt-redis/cmd"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/join", cmd.JoinUserHandler)
	auth.Post("/login", cmd.LoginUserHandler)
	auth.Post("/logout", cmd.LogoutHandler)

	todo := app.Group("/todo")
	todo.Post("/create", cmd.CreateToDo)
	todo.Post("/read", cmd.ReadToDoHandler)
	todo.Post("/update", cmd.UpdateToDoHandler)
	todo.Delete("/delete", cmd.DeleteToDoHandler)
}
