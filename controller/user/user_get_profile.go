package user

import (
	"github.com/gofiber/fiber/v2"
	"open_sea_be/helper/errorBody"
	"open_sea_be/model"
)

type GetProfileRequest struct {
	ID string `json:"id"`
}

type GetProfileResponse struct{
	User model.User `json:"user"`
}

func GetProfile(c *fiber.Ctx) error{
	getProfileRequest := GetProfileRequest{}
	getProfileRequest.ID = c.Query("user_id")


	userData, err := findUserByUserID(getProfileRequest.ID)
	if err != nil{
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	res := LoginResponse{
		User: userData,
	}

	return c.Status(200).JSON(res)
}
