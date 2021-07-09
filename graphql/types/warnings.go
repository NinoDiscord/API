package types

type Warning struct {
	GuildID string  `json:"guildID"`
	Reason  *string `json:"reason"`
	Amount  int32   `json:"amount"`
	UserID  string  `json:"userID"`
}
