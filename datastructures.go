package common

// UserDocument is the schema for information stored for a given user
type UserDocument struct {
	AccDetails AccDetailsDocument  `json:"accdetails"`
	FriendIDs  []string            `json:"friendids"`
	GamesOwned []GameOwnedDocument `json:"gamesowned"`
}

// AccDetailsDocument is the schema for information stored by
// a user for a given user (in the user collection)
type AccDetailsDocument struct {
	SteamID        string `json:"steamid"`
	Personaname    string `json:"personaname"`
	Profileurl     string `json:"profileurl"`
	Avatar         string `json:"avatar"`
	Timecreated    int    `json:"timecreated"`
	Loccountrycode string `json:"loccountrycode"`
}

// GameOwnedDocument is the schema for information stored
// by a user for a given game (in the user collection)
type GameOwnedDocument struct {
	AppID            int `json:"appid"`
	Playtime_Forever int `json:"playtime_forever"`
}

// GameInfo is the schema for information stored for each steam game
type GameInfoDocument struct {
	AppID      int    `json:"appid"`
	Name       string `json:"name"`
	ImgIconURL string `json:"imgiconurl"`
	ImgLogoURL string `json:"imglogourl"`
}

// CrawlingStatus stores the total number of friends to crawl
// for a given user and the number of profiles that have been
// crawled so far. This is used to keep track of when a given
// user has been crawled completely and processing of their
// data should start
type CrawlingStatus struct {
	TimeStarted         int64  `json:"timestarted"`
	Status              string `json:"status"`
	CrawlID             string `json:"crawlid"`
	OriginalCrawlTarget string `json:"originalcrawltarget"`
	MaxLevel            int    `json:"maxlevel"`
	TotalUsersToCrawl   int    `json:"totaluserstocrawl"`
	UsersCrawled        int    `json:"userscrawled"`
}

// UsersGraphData is the data saved for an entire crawl job that
// is later served and processed for viewing. All of a user's details
// and their friend network's details along with the ten most popular
// games of the network are saved here
type UsersGraphData struct {
	UserDetails       UsersGraphInformation   `json:"userdetails"`
	FriendDetails     []UsersGraphInformation `json:"frienddetails"`
	TopTenGameDetails []BareGameInfo          `json:"toptengamedetails"`
}

// UsersGraphInformation stores information for each user in relation
// to which user they are connected with initially in the network
type UsersGraphInformation struct {
	User         UserDocument
	FromID       string
	MaxLevel     int
	CurrentLevel int
}

// BareGameInfo correlates a steam appID to a game title
type BareGameInfo struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

// LoggingFields holds the default fields that are attached to logs
type LoggingFields struct {
	NodeName string
	NodeDC   string
	LogPaths []string
	NodeIPV4 string
	Service  string
}
