package resolvers

import "nino.sh/api/managers"

type Resolver struct {
	Db *managers.PostgresManager
}
