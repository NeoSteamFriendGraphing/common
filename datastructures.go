package common

// UserDocument is the schema for information stored for a given user
type UserDocument struct {
	SteamID            string         `json:"steamid"`
	AccDetails         Player         `json:"accDetails"`
	FriendIDs          []string       `json:"friends"`
	AmountOfGamesOwned int            `json:"amountOfGamesOwned"`
	GamesOwned         []GameInfo     `json:"gamesOwned"`
	CrawlingStatus     CrawlingStatus `json:"crawlingStatus"`
}

// GamInfo is the schema for information stored for each steam game
type GameInfo struct {
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtimeForever"`
	Playtime2Weeks  int    `json:"playtime2Weeks"`
	ImgIconURL      string `json:"imgIconUrl"`
	ImgLogoURL      string `json:"imgLogoUrl"`
}

// CrawlingStatus stores the total number of friends to crawl
// for a given user and the number of profiles that have been
// crawled so far. This is used to keep track of when a given
// user has been crawled completely and processing of their
// data should start
type CrawlingStatus struct {
	TotalUsersToCrawl int `json:"totalUsersToCrawl"`
	UsersCrawled      int `json:"usersCrawled"`
}
