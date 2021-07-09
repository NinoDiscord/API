package inputs

type AddBlacklistMetadata struct {
	Reason *string `json:"reason"`
	Issuer string  `json:"issuer"`
	Type   string  `json:"type"`
	ID     string  `json:"id"`
}
