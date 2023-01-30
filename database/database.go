package database

import (
	"context"

	"github.com/dev-hyunsang/golang-jwt-redis/config"
	"github.com/dev-hyunsang/golang-jwt-redis/ent"
	"github.com/dev-hyunsang/golang-jwt-redis/ent/user"
	"github.com/dev-hyunsang/golang-jwt-redis/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func ConnectionMySQL() (*ent.Client, error) {
	client, err := ent.Open("mysql", config.GetEnv("MYSQL_ADDR"))
	if err != nil {
		return nil, err
	}

	return client, err
}

func CreateUser(user models.User) (*ent.User, error) {
	client, err := ConnectionMySQL()
	if err != nil {
		return nil, err
	}

	data, err := client.User.
		Create().
		SetUserUUID(user.UserUUID).
		SetUserEmail(user.UserEmail).
		SetUserPassword(user.UserPassword).
		SetUserNickname(user.UserNickName).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(context.Background())

	return data, err
}

func ReadUser(email string) (*ent.User, error) {
	client, err := ConnectionMySQL()
	if err != nil {
		return nil, err
	}

	data, err := client.User.Query().
		Where(user.UserEmail(email)).
		Only(context.Background())

	return data, err
}

func UpdateUser(userData models.User) error {
	client, err := ConnectionMySQL()
	if err != nil {
		return err
	}

	_, err = client.User.Update().
		Where(user.UserUUID(userData.UserUUID)).
		SetUserEmail(userData.UserEmail).
		SetUserPassword(userData.UserPassword).
		SetUserNickname(userData.UserNickName).
		Save(context.Background())

	return err
}

func DeleteUser(userUUID uuid.UUID) error {
	client, err := ConnectionMySQL()
	if err != nil {
		return nil
	}

	user := client.User.
		Query().
		Where(user.UserUUID(userUUID)).
		OnlyX(context.Background())

	err = client.User.
		DeleteOne(user).
		Exec(context.Background())

	return err
}
