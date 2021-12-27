package model

const WalletTypeMetaMask = "METAMASK_001"
const WalletTypeTrust = "TRUSTWALLET_001"

//Illustration of wallet in cryptocurrency
//In actual world, this data is not stored in any centralized server
type Wallet struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	PublicAddress string `json:"public_address"`
	Secret        string `json:"secret"`
}
