package resolvers

func (r *Resolver) Automod() (string, error) {
	return "Hi world :3", nil
}

func (r *Resolver) ToggleShortlinks(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleBlacklist(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleMentions(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleInvites(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleDehoist(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleSpam(id string) (string, error) {
	return id, nil
}

func (r *Resolver) ToggleRaid(id string) (string, error) {
	return id, nil
}
