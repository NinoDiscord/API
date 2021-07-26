package routers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"nino.sh/api/middlewares"
	"nino.sh/api/utils"
)

func NewHealthRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.Logging)
	r.Get("/", func (w http.ResponseWriter, r *http.Request) {
		utils.SendJson(w, 200, struct { Message string }{Message: "hai!"})
	})

	return r
}
