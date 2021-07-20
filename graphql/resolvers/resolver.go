package resolvers

import (
	"nino.sh/api/controllers"
	"nino.sh/api/managers"
)

type Resolver struct {
	Controller *controllers.Controller
	Db *managers.PostgresManager
}
