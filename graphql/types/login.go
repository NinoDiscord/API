package types

type LoggedInUser struct {
	Discriminator string `json:"discriminator"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Entry         User   `json:"entry"`
	ID            string `json:"id"`
}
