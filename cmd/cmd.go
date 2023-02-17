package cmd

import (
	"fmt"
	"github.com/dev-hyunsang/golang-jwt-redis/auth"
	"log"
	"time"

	"github.com/dev-hyunsang/golang-jwt-redis/database"
	"github.com/dev-hyunsang/golang-jwt-redis/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func JoinUserHandler(c *fiber.Ctx) error {
	req := new(models.RequestJoinUser)
	if err := c.BodyParser(req); err != nil {
		log.Panic(err)
	}

	userUUID := uuid.New()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    "서버가 요청을 처리하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	// 최종적으로 DB Insert 구조체
	userData := models.User{
		UserUUID:     userUUID,
		UserEmail:    req.Email,
		UserPassword: string(hashPassword),
		UserNickName: req.NickName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result, err := database.CreateUser(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseOkJoinUser{
		Meta: models.MetaData{
			Status:     "error",
			StatusCode: fiber.StatusOK,
			Success:    true,
			Message:    "성공적으로 요청하신 유저를 만들었어요!",
		},
		Data:        *result,
		ResponsedAt: time.Now(),
	})
}

func LoginUserHandler(c *fiber.Ctx) error {
	req := new(models.RequestLoginUser)
	if err := c.BodyParser(req); err != nil {
		log.Panic(err)
	}

	// Searching User
	// 오류 발생 시 동일한 사용자 정보가 없다는 사실로 판단함.
	result, err := database.ReadUser(req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "bad request",
				StatusCode: fiber.StatusBadRequest,
				Success:    false,
				Message:    "입력해주신 정보로 사용자 정보를 찾을 수 없어요. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.UserPassword), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "bad request",
				StatusCode: fiber.StatusBadRequest,
				Success:    false,
				Message:    "입력해주신 정보로 사용자 정보를 찾을 수 없어요. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	ts, err := auth.CreateJWT(result.UserUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	err = auth.InsertRedisAuth(result.UserUUID, ts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseOkLoginUser{
		Meta: models.MetaData{
			Status:     "success",
			StatusCode: fiber.StatusOK,
			Success:    true,
			Message:    "성공적으로 로그인을 완료했어요!",
		},
		Data: models.TokenDatas{
			UserUUID:     result.UserUUID,
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		},
		ResponsedAt: time.Now(),
	})
}

func LogoutHandler(c *fiber.Ctx) error {
	au, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	deleted, err := auth.DeleteAuth(au.AccessUUID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseOkLogout{
		Meta: models.MetaData{
			Status:     "ok",
			StatusCode: fiber.StatusOK,
			Success:    true,
			Message:    "성공적으로 로그아웃 되셨습니다.",
		},
		Data:        deleted,
		ResponsedAt: time.Now(),
	})
}

func CreateToDo(c *fiber.Ctx) error {
	req := new(models.RequestCreateToDo)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "bad request",
				StatusCode: fiber.StatusBadRequest,
				Success:    false,
				Message:    "올바르지 않은 요청이예요. 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	tokenAuth, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		log.Printf("[ERROR] ExtractTokenMetaData | %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	userUUID, err := auth.FetchAuth(tokenAuth)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	parseUserUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}
	todoUUID := uuid.New()

	data := models.ToDo{
		UserUUID:    parseUserUUID,
		ToDoUUID:    todoUUID,
		ToDoTitle:   req.ToDoTitle,
		ToDoContext: req.ToDoContext,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := database.CreateToDo(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	log.Println(result)

	return c.Status(fiber.StatusOK).JSON(models.ResponseOkCreateToDo{
		Meta: models.MetaData{
			Status:     "ok",
			StatusCode: fiber.StatusOK,
			Success:    true,
			Message:    "성공적으로 새로운 할일을 생성했습니다.",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}

func UpdateToDoHandler(c *fiber.Ctx) error {
	req := new(models.RequestUpdateToDo)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusBadRequest,
				Success:    false,
				Message:    "올바르지 않은 요청이네요. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	tokenAuth, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		log.Printf("[ERROR] ExtractTokenMetaData | %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	userUUID, err := auth.FetchAuth(tokenAuth)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	parseUserUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	result, err := database.UpdateToDo(models.UpdateToDo{
		UserUUID:    parseUserUUID,
		ToDoUUID:    req.ToDoUUID,
		ToDoTitle:   req.ToDoTitle,
		ToDoContext: req.ToDoContext,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    "서버에서 처리하던 도중 알 수 없는 오류가 발생했어요. 잠시후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseOkUpdateToDo{
		Meta: models.MetaData{
			Status:     "ok",
			StatusCode: fiber.StatusOK,
			Success:    true,
			Message:    "성공적으로 할일을 수정했어요!",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}

func ReadToDoHandler(c *fiber.Ctx) error {
	tokenAuth, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		log.Printf("[ERROR] ExtractTokenMetaData | %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	userUUID, err := auth.FetchAuth(tokenAuth)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	parseUserUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	result, err := database.ReadToDo(parseUserUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseOkReadToDo{
		Meta: models.MetaData{
			Status:     "ok",
			StatusCode: fiber.StatusOK,
			Success:    true,
			Message:    "성공적으로 할일들을 불러왔습니다.",
		},
		Data:        result,
		ResponsedAt: time.Now(),
	})
}

func DeleteToDoHandler(c *fiber.Ctx) error {
	req := new(models.RequestDeleteToDo)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unprocessable entity",
				StatusCode: fiber.StatusUnprocessableEntity,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	tokenAuth, err := auth.ExtractTokenMetaData(c)
	if err != nil {
		log.Printf("[ERROR] ExtractTokenMetaData | %s", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	userUUID, err := auth.FetchAuth(tokenAuth)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "unauthorized",
				StatusCode: fiber.StatusUnauthorized,
				Success:    false,
				Message:    "잘못된 접근이예요. 로그인 이후에 시도해 주세요!",
			},
			ResponsedAt: time.Now(),
		})
	}

	parseUserUUID, err := uuid.Parse(userUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	if err := database.DeleteToDo(req.ToDoUUID, parseUserUUID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseError{
			Meta: models.MetaData{
				Status:     "bad request",
				StatusCode: fiber.StatusBadRequest,
				Success:    false,
				Message:    fmt.Sprintf("할 일을 조회할 수 없습니다. 확인 후 다시 시도해 주세요.\n%s", err.Error()),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseOk{
		Meta: models.MetaData{
			Status:     "ok",
			StatusCode: fiber.StatusOK,
			Success:    true,
			Message:    "정상적으로 할일을 삭제했습니다.",
		},
		ResponsedAt: time.Now(),
	})
}
