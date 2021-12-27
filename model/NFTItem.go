package model

const NFTCategoryArt = "Art"
const NFTCategoryMusic = "Music"
const NFTCategoryVirtualWorld = "VirtualWorld"
const NFTCategoryRealWorldOwnership = "RWO"

const VerifiedUserNFTTrue = 1
const VerifiedUserNFTFalse = 0

type NFTItem struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryID   string `json:"category_id"`
	Price        int    `json:"price"`
	Owner        string `json:"owner"`
	VerifiedUser int    `json:"verified_user"`
}
