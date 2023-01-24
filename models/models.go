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

type RequestJoinUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
}

type RequestLoginUser struct {
	Email string `json:"email"`
}

type MetaData struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

type ResponseOkJoinUser struct {
	Meta        MetaData  `json:"meta"`
	Data        ent.User  `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ResponseError struct {
	Meta        MetaData  `json:"meta"`
	ResponsedAt time.Time `json:"responsed_at"`
}
