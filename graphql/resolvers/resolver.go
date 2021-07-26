package resolvers

import (
	"nino.sh/api/controllers"
	"nino.sh/api/managers"
	"nino.sh/api/redis"
)

type Resolver struct {
	Controller *controllers.Controller
	Redis *redis.Redis
	Db *managers.PostgresManager
}
