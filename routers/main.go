package routers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"nino.sh/api/utils"
)

type MainResponse struct {
	Message string `json:"message"`
	Docs    string `json:"docs_url"`
}

func NewMainRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func (w http.ResponseWriter, r *http.Request) {
		utils.SendJson(w, 200, &MainResponse{
			Message: "hello world :D",
			Docs: "https://nino.sh/docs",
		})
	})

	return router
}
