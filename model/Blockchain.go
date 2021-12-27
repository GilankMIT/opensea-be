package model

type Blockchain struct {
	ID              string `json:"id"`
	Amount          int64  `json:"amount"`
	NFTItemID       int    `json:"nft_item_id"`
	PublicAddress   string `json:"public_address"`
	PreviousAddress string `json:"previous_address"`
	Signature       string `json:"signature"`
}
