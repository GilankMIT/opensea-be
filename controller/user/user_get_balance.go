package user

import (
	"github.com/gofiber/fiber/v2"
)

type GetBalanceRequest struct {
	UserID string `json:"user_id"`
}

type GetBalanceResponse struct {
	Balance      int `json:"balance"`
	LinkedWallet int `json:"linked_wallet"`
}

func GetBalance(c *fiber.Ctx) error {
	return nil
}
