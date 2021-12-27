package user

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"open_sea_be/helper/errorBody"
	"open_sea_be/model"
)

type RegisterConsultRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RegisterConsultResponse struct{
	User model.User `json:"user"`
}

func UserRegisterConsult(c *fiber.Ctx) error{
	registerConsultRequest := RegisterConsultRequest{}
	if err := c.BodyParser(&registerConsultRequest); err != nil {
		fmt.Println("error = ",err)
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	if len(registerConsultRequest.Password) < 6{
		return c.Status(400).JSON(errorBody.ErrorBody{Error: "password length cannot be les than 6"})
	}

	_, err := findUserByEmail(registerConsultRequest.Email)
	if err == nil{
		return c.Status(400).JSON(errorBody.ErrorBody{Error: errors.New("user with that email already exist").Error()})
	}

	res := RegisterConsultResponse{User: model.User{}}

	return c.Status(200).JSON(res)
}
