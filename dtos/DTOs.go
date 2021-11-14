package dtos

import (
	"github.com/neosteamfriendgraphing/common"
)

// SaveUserDTO is the input schema for saving users to the database. It takes
// the original crawl target user (that initally caused this crawl) and the
// current user to be saved
type SaveUserDTO struct {
	OriginalCrawlTarget string              `json:"orginalcrawltarget"`
	CurrentLevel        int                 `json:"currentlevel"`
	MaxLevel            int                 `json:"maxlevel"`
	User                common.UserDocument `json:"user"`
	GamesOwnedFull      []common.GameInfo   `json:"gamesownedfull"`
}

// GetUserDTO is the returned data when user has been successfully
// found in the database
type GetUserDTO struct {
	Status string              `json:"status"`
	User   common.UserDocument `json:"user"`
}
