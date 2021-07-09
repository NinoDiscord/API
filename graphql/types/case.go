package types

type Case struct {
	Attachments []string `json:"attachments"`
	ModeratorID string   `json:"moderatorID"`
	MessageID   *string  `json:"messageID"`
	VictimID    string   `json:"victimID"`
	Reason      *string  `json:"reason"`
	Type        []string `json:"type"`
	Time        *int32   `json:"time"`
}
