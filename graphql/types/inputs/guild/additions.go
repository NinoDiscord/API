package guild

type AddPunishmentMetadata struct {
	Warnings int32  `json:"warnings"`
	Soft     bool   `json:"soft"`
	Time     *int32  `json:"time"`
	Type     string `json:"type"`
	ID       string  `json:"id"`
}

type AddWarningMetadata struct {
	Reason *string `json:"reason"`
	Amount int32  `json:"amount"`
	Victim string `json:"victim"`
	ID     string `json:"id"`
}
