package common

type UserDocument struct {
	SteamID            string         `json:"steamid"`
	AccDetails         Player         `json:"accDetails"`
	FriendIDs          []string       `json:"friends"`
	AmountOfGamesOwned int            `json:"amountOfGamesOwned"`
	GamesOwned         []GameInfo     `json:"gamesOwned"`
	CrawlingStatus     CrawlingStatus `json:"crawlingStatus"`
}

type GameInfo struct {
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtimeForever"`
	Playtime2Weeks  int    `json:"playtime2Weeks"`
	ImgIconURL      string `json:"imgIconUrl"`
	ImgLogoURL      string `json:"imgLogoUrl"`
}

type CrawlingStatus struct {
	TotalUsersToCrawl int `json:"totalUsersToCrawl"`
	UsersCrawled      int `json:"usersCrawled"`
}
