package nft_item

import (
	"github.com/gofiber/fiber/v2"
	"open_sea_be/model"
)

type GetAllResponse struct{
	NFTItems []model.NFTItem `json:"nft_items"`
}

func GetAll(c *fiber.Ctx) error{
	res := GetAllResponse{NFTItems: items}
	return c.Status(200).JSON(res)
}
