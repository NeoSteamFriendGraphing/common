package common

// SteamAPIResponse is the response from the steam web API
// for /getPlayerSummary calls
type SteamAPIResponse struct {
	Response Response `json:"response"`
}

// Response is filler that the steam web API includes
type Response struct {
	Players []Player `json:"players"`
}

// Player holds all details for a given user returned by the steam web API for
// the /getPlayerSummary endpoint
type Player struct {
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Commentpermission        int    `json:"commentpermission"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname"`
	Primaryclanid            string `json:"primaryclanid"`
	Timecreated              int    `json:"timecreated"`
	Personastateflags        int    `json:"personastateflags"`
	Loccountrycode           string `json:"loccountrycode"`
}

type UserDetails struct {
	SteamID int64       `json:"steamID"`
	Friends Friendslist `json:"friendsList"`
}

// FriendsList holds all friends for a given user
type Friendslist struct {
	Friends []Friend `json:"friends"`
}

// Friend is profile information returned for each friend by
// the steam web API for /getFriendsList calls
type Friend struct {
	Username     string `json:"username"`
	Steamid      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

// GamesOwnedSteamResponse is the response from the steam web API
// for /getOwnedGames calls
type GamesOwnedSteamResponse struct {
	Response GamesOwnedResponse `json:"response"`
}

// GamesOwnedResponse is filler returned by the steam web API
type GamesOwnedResponse struct {
	GameCount int    `json:"game_count"`
	Games     []Game `json:"games"`
}

// Game is returned by the steam web API for each game
// from calls to /getOwnedGames
type Game struct {
	Appid                    int    `json:"appid"`
	Name                     string `json:"name"`
	PlaytimeForever          int    `json:"playtime_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	ImgLogoURL               string `json:"img_logo_url"`
	PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
	PlaytimeMacForever       int    `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats,omitempty"`
	Playtime2Weeks           int    `json:"playtime_2weeks,omitempty"`
}
