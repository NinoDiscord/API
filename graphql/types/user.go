package types

type User struct {
	Language string   `json:"language"`
	Prefixes []string `json:"prefixes"`
}
