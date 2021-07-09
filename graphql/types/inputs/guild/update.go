package guild

type UpdateGuildMetadata struct {
	ModLogChannelID *string `json:"modLogChannelID"`
	MutedRoleID     *string `json:"mutedRoleID"`
	Language        *string `json:"language"`
	ID              string  `json:"id"`
}

type UpdateLoggingMetadata struct {
	ChannelID *string `json:"channelID"`
	ID        string  `json:"id"`
}
