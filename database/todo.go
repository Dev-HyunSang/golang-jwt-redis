package database

import (
	"context"
	"github.com/dev-hyunsang/golang-jwt-redis/ent"
	"github.com/dev-hyunsang/golang-jwt-redis/models"
)

func CreateToDo(todo *models.ToDo) (*ent.ToDo, error) {
	client, err := ConnectionMySQL()
	if err != nil {
		return nil, err
	}

	result, err := client.ToDo.Create().
		SetTodoUUID(todo.ToDoUUID).
		SetUserUUID(todo.UserUUID).
		SetTodoTitle(todo.ToDoTitle).
		SetTodoContext(todo.ToDoContext).
		SetCratedAt(todo.CreatedAt).
		SetUpdatedAt(todo.UpdatedAt).
		Save(context.Background())

	return result, err
}
