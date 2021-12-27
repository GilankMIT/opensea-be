package user

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"open_sea_be/helper/errorBody"
	"open_sea_be/model"
)

type WalletLinkageRequest struct {
	UserID          string `json:"user_id"`
	WalletPublicKey string `json:"wallet_public_key"`
	Secret          string `json:"secret"`
}

type WalletLinkageResponse struct {
	Message string `json:"message"`
	Balance int `json:"balance"`
}

func UserWalletLinkage(c *fiber.Ctx) error {
	walletLinkageRequest := WalletLinkageRequest{}
	if err := c.BodyParser(&walletLinkageRequest); err != nil {
		fmt.Println("error = ", err)
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	walletData, err := findWalletByPublicKeyAndSecret(walletLinkageRequest.WalletPublicKey, walletLinkageRequest.Secret)
	if err != nil {
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	//link user to wallet
	err = linkWalletToUser(walletLinkageRequest.UserID, walletData)
	if err != nil {
		return c.Status(400).JSON(errorBody.ErrorBody{Error: err.Error()})
	}

	res := WalletLinkageResponse{Message: "wallet linked successfully", Balance: getUserBalanceFromWallet(walletData.PublicAddress)}
	return c.Status(200).JSON(res)
}

func findWalletByPublicKeyAndSecret(publicKey, secret string) (model.Wallet, error) {
	for _, walletData := range reservedWallets {
		if walletData.PublicAddress == publicKey && walletData.Secret == walletData.Secret {
			return walletData, nil
		}
	}

	return model.Wallet{}, errors.New("cannot get wallet")
}

func linkWalletToUser(userId string, walletData model.Wallet) error {
	//find user by user id
	userData, err := findUserByUserID(userId)
	if err != nil {
		return err
	}

	userData.Wallets = append(userData.Wallets, walletData)

	//update user data list
	alteredIdx := 0
	for idx, userData := range userDataList {
		if userData.ID == userId {
			alteredIdx = idx
		}
	}

	userDataList[alteredIdx] = userData

	return nil
}

func findUserByUserID(userID string) (model.User, error) {
	for _, userData := range userDataList {
		if userData.ID == userID {
			return userData, nil
		}
	}

	return model.User{}, errors.New("cannot find user by the user id")
}

func getUserBalanceFromWallet(walletPublicKey string) int{
	return 10000 //TODO : Please continue this!
}
