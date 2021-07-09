package types

type Logging struct {
	IgnoreChannels []string `json:"ignoreChannels"`
	IgnoreUsers    []string `json:"ignoreUsers"`
	ChannelID      *string  `json:"channelID"`
	Enabled 	   bool     `json:"enabled"`
	Events         []string `json:"events"`
}
