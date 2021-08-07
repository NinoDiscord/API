package types

type Automod struct {
	WhitelistChannels []string `json:"whitelist_channels"`
	BlacklistedWords  []string `json:"blacklisted_words"`
	ShortLinks        bool `json:"shortLinks"`
	Blacklist         bool `json:"blacklist"`
	Mentions          bool `json:"mentions"`
	GuildId           string `json:"guild_id"`
	Invites           bool `json:"invites"`
	Dehoist           bool `json:"dehoist"`
	Spam              bool `json:"spam"`
	Raid 	          bool `json:"raid"`
}
