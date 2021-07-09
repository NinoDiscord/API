package routers

import "github.com/go-chi/chi/v5"

func NewHealthRouter() chi.Router {
	r := chi.NewRouter()
	return r
}
