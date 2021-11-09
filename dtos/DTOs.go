package dtos

import (
	"github.com/neosteamfriendgraphing/common"
)

// SaveUserDTO is the input schema for saving users to the database. It takes
// the original crawl target user (that initally caused this crawl) and the
// current user to be saved
type SaveUserDTO struct {
	OriginalCrawlTarget string              `json:"orginalcrawltarget"`
	User                common.UserDocument `json:"user"`
}
