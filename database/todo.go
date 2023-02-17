package database

import (
	"context"
	"github.com/dev-hyunsang/golang-jwt-redis/ent"
	"github.com/dev-hyunsang/golang-jwt-redis/ent/todo"
	"github.com/dev-hyunsang/golang-jwt-redis/models"
	"github.com/google/uuid"
	"time"
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

func UpdateToDo(updateToDo models.UpdateToDo) (int, error) {
	client, err := ConnectionMySQL()
	if err != nil {
		return 0, err
	}

	// 유효성 검사 및 제목, 내용 업데이트
	result, err := client.ToDo.Update().
		Where(todo.TodoUUID(updateToDo.ToDoUUID)).
		Where(todo.UserUUID(updateToDo.UserUUID)).
		SetTodoTitle(updateToDo.ToDoTitle).
		SetTodoContext(updateToDo.ToDoContext).
		SetUpdatedAt(time.Now()).
		Save(context.Background())

	return result, err
}

func ReadToDo(userUUID uuid.UUID) ([]*ent.ToDo, error) {
	client, err := ConnectionMySQL()
	if err != nil {
		return nil, err
	}

	result, err := client.ToDo.Query().
		Where(todo.UserUUID(userUUID)).
		All(context.Background())

	return result, err
}

func DeleteToDo(todoUUID, userUUID uuid.UUID) error {
	client, err := ConnectionMySQL()
	if err != nil {
		return err
	}

	todoResult := client.ToDo.Query().
		Where(todo.TodoUUID(todoUUID)).
		Where(todo.UserUUID(userUUID)).
		OnlyX(context.Background())

	err = client.ToDo.DeleteOne(todoResult).
		Exec(context.Background())

	return err
}
