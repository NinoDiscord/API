package graphql

import (
	"encoding/json"
	"github.com/graph-gophers/graphql-go"
	"io/ioutil"
	"net/http"
	"nino.sh/api/controllers"
	"nino.sh/api/graphql/resolvers"
	"nino.sh/api/managers"
)

type Manager struct {
	Postgres *managers.PostgresManager
	Schema *graphql.Schema
}

type Body struct {
	Variables     map[string]interface{} `json:"variables"`
	OperationName string `json:"operationName"`
	Query         string `json:"query"`
}

func NewGraphQLManager(postgres *managers.PostgresManager) *Manager {
	return &Manager{
		Postgres: postgres,
		Schema: nil,
	}
}

func (gql *Manager) GenerateSchema() error {
	contents, err := ioutil.ReadFile("./schema.gql"); if err != nil {
		return err
	}

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	items := string(contents)
	schema := graphql.MustParseSchema(items, &resolvers.Resolver{
		Db: gql.Postgres,
		Controller: &controllers.Controller{},
	}, opts...)

	gql.Schema = schema
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
