package user

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"open_sea_be/helper/errorBody"
	"open_sea_be/model"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct{
	User model.User `json:"user"`
}

func UserLogin(c *fiber.Ctx) error{
	loginRequest := LoginRequest{}
	if err := c.BodyParser(&loginRequest); err != nil {
		fmt.Println("error = ",err)
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	userData, err := findUserByEmail(loginRequest.Email)
	if err != nil{
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginRequest.Password)); err != nil{
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	res := LoginResponse{
		User: userData,
	}

	return c.Status(200).JSON(res)
}


func findUserByEmail(email string) (model.User, error){
	for _, userData := range userDataList{
		if userData.Email == email{
			return userData, nil
		}
	}

	return model.User{}, errors.New("cannot find user")
}
