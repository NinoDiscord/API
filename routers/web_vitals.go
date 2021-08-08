package routers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"nino.sh/api/metrics"
	"nino.sh/api/utils"
)

type WebVitalBody struct {
	Name string
	Value float32 // im going to assume it's float32
}

func NewWebVitalsRouter(m *metrics.Metrics) chi.Router {
	router := chi.NewRouter()
	router.Get("/", func (w http.ResponseWriter, req *http.Request) {
		utils.SendJson(w, 405, struct{ Message string }{
			Message: "/web-vitals only supports POST requests.",
		})
	})

	router.Post("/", func (w http.ResponseWriter, req *http.Request) {
		var body WebVitalBody
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			utils.SendJson(w, 400, struct { Message string }{
				Message: "unable to serialize body",
			})
		}

		// i should add authentication but oh well
		m.SetWebVital(body.Name, body.Value)
	})

	return router
}
