package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"open_sea_be/controller/nft_item"
	"open_sea_be/controller/user"
)

func main(){
	app := fiber.New()

	app.Post("/api/register/consult", func(c *fiber.Ctx) error {
		return user.UserRegisterConsult(c)
	})

	app.Post("/api/register", func(c *fiber.Ctx) error {
		return user.UserRegister(c)
	})

	app.Post("/api/login", func(c *fiber.Ctx) error{
		return user.UserLogin(c)
	})

	app.Post("/api/link-wallet", func(c *fiber.Ctx) error {
		return user.UserWalletLinkage(c)
	})

	app.Get("/api/profile", func(c *fiber.Ctx) error{
		return user.GetProfile(c)
	})

	app.Get("/api/my-balance", func(c *fiber.Ctx) error{
		return user.GetProfile(c)
	})

	app.Get("/api/nft_items", func(c *fiber.Ctx) error{
		return nft_item.GetAll(c)
	})

	app.Get("/api/nft_items/top", func(c *fiber.Ctx) error{
		return nft_item.GetTopItems(c)
	})

	app.Get("/api/nft_items/detail", func(c *fiber.Ctx) error{
		return nft_item.GetItemDetail(c)
	})

	log.Println("app listening to 8010")
	if err := app.Listen("192.168.0.16:8010"); err != nil{
		log.Fatal(err.Error())
	}
}