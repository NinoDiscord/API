package graphql

import (
	"encoding/json"
	"github.com/graph-gophers/graphql-go"
	"io/ioutil"
	"net/http"
)

type Resolver struct {}

type Manager struct {
	Schema *graphql.Schema
}

type Body struct {
	OperationName string 				 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
	Query         string  			     `json:"query"`
}

func NewGraphQLManager() *Manager {
	return &Manager{
		Schema: nil,
	}
}

func (gql *Manager) GenerateSchema() error {
	contents, err := ioutil.ReadFile("./schema.gql"); if err != nil {
		return err
	}

	_ = string(contents)
	return nil
}

func (gql *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params Body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	res := gql.Schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
	data, err := json.Marshal(res); if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_, _ = w.Write(data)
}
