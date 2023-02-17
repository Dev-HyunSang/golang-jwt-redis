package models

import (
	"time"

	"github.com/dev-hyunsang/golang-jwt-redis/ent"
	"github.com/google/uuid"
)

type User struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"user_password"`
	UserNickName string    `json:"user_nickname"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TokenDetails struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	AccessUUID   uuid.UUID `json:"access_uuid"`
	RefreshUUID  uuid.UUID `json:"refresh_uuid"`
	AtExpires    int64     `json:"at_expires"`
	RtExpires    int64     `json:"rt_expires"`
}

type AccessDetails struct {
	AccessUUID string `json:"access_uuid"`
	UserUUID   string `json:"user_uuid"`
}

type ToDo struct {
	UserUUID    uuid.UUID `json:"user_uuid"`
	ToDoUUID    uuid.UUID `json:"todo_uuid"`
	ToDoTitle   string    `json:"todo_title"`
	ToDoContext string    `json:"todo_context"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
type UpdateToDo struct {
	UserUUID    uuid.UUID `json:"user_uuid"`
	ToDoUUID    uuid.UUID `json:"todo_uuid"`
	ToDoTitle   string    `json:"todo_title"`
	ToDoContext string    `json:"todo_context"`
}

// === Request ===
type RequestJoinUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
}

type RequestLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestCreateToDo struct {
	ToDoTitle   string `json:"todo_title"`
	ToDoContext string `json:"todo_context"`
}

type RequestUpdateToDo struct {
	ToDoUUID    uuid.UUID `json:"todo_uuid"`
	ToDoTitle   string    `json:"todo_title"`
	ToDoContext string    `json:"todo_context"`
}

type RequestDeleteToDo struct {
	ToDoUUID uuid.UUID `json:"todo_uuid"`
}

// === Response ===
type MetaData struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

type ResponseError struct {
	Meta        MetaData  `json:"meta"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseOkJoinUser struct {
	Meta        MetaData  `json:"meta"`
	Data        ent.User  `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type TokenDatas struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
type ResponseOk struct {
	Meta        MetaData  `json:"meta"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseOkLoginUser struct {
	Meta        MetaData   `json:"meta"`
	Data        TokenDatas `json:"data"`
	ResponsedAt time.Time  `json:"responsed_at"`
}

type ResponseOkCreateToDo struct {
	Meta        MetaData  `json:"meta"`
	Data        *ent.ToDo `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseOkUpdateToDo struct {
	Meta        MetaData  `json:"meta"`
	Data        int       `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseOkReadToDo struct {
	Meta        MetaData    `json:"meta"`
	Data        []*ent.ToDo `json:"data"`
	ResponsedAt time.Time   `json:"responsed_at"`
}

type ResponseOkLogout struct {
	Meta        MetaData  `json:"meta"`
	Data        int64     `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}
