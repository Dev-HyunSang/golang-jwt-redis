// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dev-hyunsang/golang-jwt-redis/ent/schema"
	"github.com/dev-hyunsang/golang-jwt-redis/ent/todo"
	"github.com/dev-hyunsang/golang-jwt-redis/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.ToDo{}.Fields()
	_ = todoFields
	// todoDescTodoUUID is the schema descriptor for todo_uuid field.
	todoDescTodoUUID := todoFields[0].Descriptor()
	// todo.DefaultTodoUUID holds the default value on creation for the todo_uuid field.
	todo.DefaultTodoUUID = todoDescTodoUUID.Default.(func() uuid.UUID)
	// todoDescUserUUID is the schema descriptor for user_uuid field.
	todoDescUserUUID := todoFields[1].Descriptor()
	// todo.DefaultUserUUID holds the default value on creation for the user_uuid field.
	todo.DefaultUserUUID = todoDescUserUUID.Default.(func() uuid.UUID)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoFields[4].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	// todoDescCratedAt is the schema descriptor for crated_at field.
	todoDescCratedAt := todoFields[5].Descriptor()
	// todo.DefaultCratedAt holds the default value on creation for the crated_at field.
	todo.DefaultCratedAt = todoDescCratedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserUUID is the schema descriptor for user_uuid field.
	userDescUserUUID := userFields[0].Descriptor()
	// user.DefaultUserUUID holds the default value on creation for the user_uuid field.
	user.DefaultUserUUID = userDescUserUUID.Default.(func() uuid.UUID)
	// userDescUserEmail is the schema descriptor for user_email field.
	userDescUserEmail := userFields[1].Descriptor()
	// user.DefaultUserEmail holds the default value on creation for the user_email field.
	user.DefaultUserEmail = userDescUserEmail.Default.(string)
	// userDescUserPassword is the schema descriptor for user_password field.
	userDescUserPassword := userFields[2].Descriptor()
	// user.DefaultUserPassword holds the default value on creation for the user_password field.
	user.DefaultUserPassword = userDescUserPassword.Default.(string)
	// userDescUserNickname is the schema descriptor for user_nickname field.
	userDescUserNickname := userFields[3].Descriptor()
	// user.DefaultUserNickname holds the default value on creation for the user_nickname field.
	user.DefaultUserNickname = userDescUserNickname.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
}