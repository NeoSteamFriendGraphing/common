package common

// UserDocument is the schema for information stored for a given user
type UserDocument struct {
	SteamID    string     `json:"steamid"`
	AccDetails Player     `json:"accdetails"`
	FriendIDs  []string   `json:"friendids"`
	GamesOwned []GameInfo `json:"gamesowned"`
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
	OriginalCrawlTarget string `json:"originalcrawltarget"`
	MaxLevel            int    `json:"maxlevel"`
	TotalUsersToCrawl   int    `json:"totaluserstocrawl"`
	UsersCrawled        int    `json:"userscrawled"`
}

// LoggingFields holds the default fields attached to logs
type LoggingFields struct {
	NodeName string
	NodeDC   string
	LogPaths []string
	NodeIPV4 string
}
