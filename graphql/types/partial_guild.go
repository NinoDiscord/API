package types

type PartialGuild struct {
	Permissions   string
	Features      []string
	Owner         bool
	Icon          *string
	Name          string
	ID            string
}
