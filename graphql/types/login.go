package types

type LoggedInUser struct {
	Discriminator string  `json:"discriminator"`
	Username      string  `json:"username"`
	Avatar        *string `json:"avatar"`
	Guilds        []PartialGuild `json:"guilds"`
	Entry         User    `json:"entry"`
	ID            string  `json:"id"`
}
