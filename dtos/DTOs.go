package dtos

import (
	"github.com/neosteamfriendgraphing/common"
)

// SaveUserDTO is the input schema for saving users to the database. It takes
// the original crawl target user (that initally caused this crawl) and the
// current user to be saved
type SaveUserDTO struct {
	OriginalCrawlTarget string                    `json:"orginalcrawltarget"`
	CrawlID             string                    `json:"crawlid"`
	CurrentLevel        int                       `json:"currentlevel"`
	MaxLevel            int                       `json:"maxlevel"`
	User                common.UserDocument       `json:"user"`
	GamesOwnedFull      []common.GameInfoDocument `json:"gamesownedfull"`
}

// GetUserDTO is the returned data when user has been successfully
// found in the database
type GetUserDTO struct {
	Status string              `json:"status"`
	User   common.UserDocument `json:"user"`
}

// SaveCrawlingStatsDTO is the input struct used when saving just
// the crawling status. This is used when a user is already found
// in the database an no steam API calls are needed before updating
// the crawling status
type SaveCrawlingStatsDTO struct {
	CurrentLevel   int                   `json:"currentlevel"`
	CrawlingStatus common.CrawlingStatus `json:"crawlingstatus"`
}

// GetCrawlingStatusDTO is returned from GET /getcrawlingstatus
type GetCrawlingStatusDTO struct {
	Status         string                `json:"status"`
	CrawlingStatus common.CrawlingStatus `json:"crawlingstatus"`
}

// GetGraphableDataForUserDTO is returned from GET /getgraphabledata
type GetGraphableDataForUserDTO struct {
	Username  string   `json:"username"`
	SteamID   string   `json:"steamid"`
	FriendIDs []string `json:"friendids"`
}

// GetUsernamesFromSteamIDsInputDTO is the input body accepted to
// POST /getusernamesfromsteamids
type GetUsernamesFromSteamIDsInputDTO struct {
	SteamIDs []string `json:"steamids"`
}

// GetUsernamesFromSteamIDsDTO is returned from calls to
// POST /getusernamesfromsteamids
type GetUsernamesFromSteamIDsDTO struct {
	SteamIDAndUsername []SteamIDAndUsername `json:"steamidsandusername"`
}

type SteamIDAndUsername struct {
	SteamID  string `json:"steamid"`
	Username string `json:"username"`
}
