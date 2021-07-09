package types

type Guild struct {
	ModLogChannelID *string  `json:"modlogChannelID"`
	MutedRoleID     *string  `json:"mutedRoleID"`
	Prefixes        []string `json:"prefixes"`
	Language        string   `json:"language"`
}
