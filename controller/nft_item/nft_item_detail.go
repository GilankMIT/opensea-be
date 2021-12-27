package nft_item

import (
	"github.com/gofiber/fiber/v2"
	"open_sea_be/helper/errorBody"
	"open_sea_be/model"
)

type  GetNFTItemDetailRequest struct{
	ID string `json:"id"`
}

type  GetNFTItemDetailResponse struct{
	NFTItem model.NFTItem `json:"nft_item"`
}

func GetItemDetail(c *fiber.Ctx) error{
	getProfileRequest := GetNFTItemDetailRequest{}
	getProfileRequest.ID = c.Query("item_id")
	res := GetNFTItemDetailResponse{}

	for _, nftItemNode := range items{
		if nftItemNode.ID == getProfileRequest.ID{
			res.NFTItem = nftItemNode
			return c.Status(200).JSON(res)
		}
	}

	return c.Status(400).JSON(errorBody.ErrorBody{Error: "item not found"})
}
