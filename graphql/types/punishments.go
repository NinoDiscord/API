package types

type Punishment struct {
	Warnings int32    `json:"warnings"`
	Index    int32    `json:"index"`
	Soft     bool     `json:"soft"`
	Time     *int32   `json:"time"`
	Type     []string `json:"type"`
}
