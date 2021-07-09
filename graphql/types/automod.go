package types

type Automod struct {
	ShortLinks bool `json:"shortLinks"`
	Blacklist  bool `json:"blacklist"`
	Mentions   bool `json:"mentions"`
	Invites    bool `json:"invites"`
	Dehoist    bool `json:"dehoist"`
	Spam       bool `json:"spam"`
	Raid 	   bool `json:"raid"`
}
