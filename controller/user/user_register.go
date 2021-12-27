package user

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"open_sea_be/helper/errorBody"
	"open_sea_be/model"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RegisterResponse struct{
	User model.User `json:"user"`
}

func UserRegister(c *fiber.Ctx) error{
	registerRequest := RegisterRequest{}
	if err := c.BodyParser(&registerRequest); err != nil {
		fmt.Println("error = ",err)
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	_, err := findUserByEmail(registerRequest.Email)
	if err == nil{
		return c.Status(400).JSON(errorBody.ErrorBody{Error: errors.New("user with that email already exist").Error()})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 12)

	newUser := model.User{
		ID:       uuid.NewV4().String(),
		Email:    registerRequest.Email,
		Password: string(password),
		Username: registerRequest.Username,
	}

	userDataList = append(userDataList, newUser)

	return c.Status(200).JSON(newUser)
}
