package routers

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"nino.sh/api/graphql"
	"nino.sh/api/middlewares"
	"nino.sh/api/utils"
	"os"
)

func NewGraphQLRouter(handler *graphql.Manager) chi.Router {
	router := chi.NewRouter()

	router.Use(middlewares.Logging)
	router.Post("/", handler.ServeHTTP)
	router.Get("/", func (w http.ResponseWriter, r *http.Request) {
		if os.Getenv("GO_ENV") == "development" {
			t := template.New("Nino: GraphQL Playground")
			t, err := t.Parse(utils.PlaygroundTemplate); if err != nil {
				http.Error(w, err.Error(), 500)
			}

			data := utils.PlaygroundTemplateData{Endpoint: "http://localhost:6645/graphql"}
			if err := t.ExecuteTemplate(w, "index", data); err != nil {
				http.Error(w, err.Error(), 500)
			}

			return
		}

		utils.SendJson(w, 404, utils.Response{
			Message: "Cannot GET /graphql",
		})
	})

	return router
}
