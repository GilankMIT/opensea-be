package nft_item

import (
	"github.com/gofiber/fiber/v2"
	"open_sea_be/model"
)

type GetTopResponse struct{
	NFTItems []model.NFTItem `json:"nft_items"`
}

func GetTopItems(c *fiber.Ctx) error{
	res := GetTopResponse{NFTItems: topItems}
	return c.Status(200).JSON(res)
}
