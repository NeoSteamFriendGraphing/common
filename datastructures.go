package common

// UserDocument is the schema for information stored for a given user
type UserDocument struct {
	SteamID            string         `json:"steamid"`
	AccDetails         Player         `json:"accdetails"`
	FriendIDs          []string       `json:"friendids"`
	AmountOfGamesOwned int            `json:"amountofgamesowned"`
	GamesOwned         []GameInfo     `json:"gamesowned"`
	CrawlingStatus     CrawlingStatus `json:"crawlingstatus"`
}

// GamInfo is the schema for information stored for each steam game
type GameInfo struct {
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtimeforever"`
	Playtime2Weeks  int    `json:"playtime2weeks"`
	ImgIconURL      string `json:"imgiconurl"`
	ImgLogoURL      string `json:"imglogourl"`
}

// CrawlingStatus stores the total number of friends to crawl
// for a given user and the number of profiles that have been
// crawled so far. This is used to keep track of when a given
// user has been crawled completely and processing of their
// data should start
type CrawlingStatus struct {
	TotalUsersToCrawl int `json:"totaluserstocrawl"`
	UsersCrawled      int `json:"userscrawled"`
}
