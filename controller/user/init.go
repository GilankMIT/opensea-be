package user

import "open_sea_be/model"

var reservedWalletReneLiu = []model.Wallet{
	{
		ID:            "1188b6c7-677f-486b-8e0b-65ceab484f9d",
		Type:          model.WalletTypeMetaMask,
		PublicAddress: "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvJuQkkl+fyb5DwPRTubbxcouewILyBqPn66kZchT7Y2eWJMLUjobu8xAh7jtb3PvO1HgNgeTOQDxJsmMD6pUmTvizDKF3rZ5NFmWicP9yl2VGPo1yGTE+EmhChH+vB0tUNqJeaMjsbGi1QxKjEX8g3W7DaPBWboL5sYICDz8eTQIDAQAB",
		Secret:        "root001",
	},
}

var reservedWallets = []model.Wallet{
	{
		ID:            "1188b6c7-677f-486b-8e0b-65ceab484f01",
		Type:          model.WalletTypeMetaMask,
		PublicAddress: "11GfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvJuQkkl+fyb5DwPRTubbxcouewILyBqPn66kZchT7Y2eWJMLUjobu8xAh7jtb3PvO1HgNgeTOQDxJsmMD6pUmTvizDKF3rZ5NFmWicP9yl2VGPo1yGTE+EmhChH+vB0tUNqJeaMjsbGi1QxKjEX8g3W7DaPBWboL5sYICDz8eTQIDAQAB",
		Secret:        "root002",
	},

	{
		ID:            "1088b6c7-677f-486b-8e0b-65ceab484f01",
		Type:          model.WalletTypeTrust,
		PublicAddress: "10GfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvJuQk01+fyb5DwPRTubbxcouewILyBqPn66kZchT7Y2eWJMLUjobu8xAh7jtb3PvO1HgNgeTOQDxJsmMD6pUmTvizDKF3rZ5NFmWicP9yl2VGPo1yGTE+EmhChH+vB0tUNqJeaMjsbGi1QxKjEX8g3W7DaPBWboL5sYICDz8eTQIDAQAB",
		Secret:        "root003",
	},

	{
		ID:            "1088b6c7-677f-486b-8e0b-65ceab484f02",
		Type:          model.WalletTypeTrust,
		PublicAddress: "10GfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvJuQk02+fyb5DwPRTubbxcouewILyBqPn66kZchT7Y2eWJMLUjobu8xAh7jtb3PvO1HgNgeTOQDxJsmMD6pUmTvizDKF3rZ5NFmWicP9yl2VGPo1yGTE+EmhChH+vB0tUNqJeaMjsbGi1QxKjEX8g3W7DaPBWboL5sYICDz8eTQIDAQAB",
		Secret:        "root003",
	},
}

var reservedUser = []model.User{
	{
		ID:       "1b22a0af-4d26-415d-8249-3578ae9bd4b6",
		Email:    "reneliu@mail.cn",
		Password: "$2a$12$Nws0LuyOHwvlX4UepUD13.yqWSA7H2tRIj8wZEdfiJFRDFVLRti3y",
		Username: "rene.liu刘若英",
		Location: "Shenzen, China",
		Wallets:  reservedWalletReneLiu,
	},
}

var userDataList = reservedUser
