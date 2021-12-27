package model

type User struct{
	ID string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Location string `json:"location"`
	Wallets []Wallet `json:"wallets"`
}
