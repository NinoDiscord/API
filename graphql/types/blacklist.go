package types

type Blacklist struct {
	Reason *string  `json:"reason"`
	Issuer string   `json:"issuer"`
	Type   []string `json:"type"`
}
