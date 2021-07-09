package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"nino.sh/api/graphql"
	"nino.sh/api/managers"
	"nino.sh/api/routers"
	"os"
)

func main() {
	err := godotenv.Load(".env"); if err != nil {
		panic(err)
	}

	logrus.Info("Now bootstrapping API...")
	postgres := managers.NewPostgresManager()
	err = postgres.GetConnection(); if err != nil {
		logrus.Fatalf("Unable to connect to Postgres: %o", err)
		os.Exit(1)
	}

	gql := graphql.NewGraphQLManager()
	if err := gql.GenerateSchema(); err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Mount("/", routers.NewMainRouter())
	router.Mount("/health", routers.NewHealthRouter())
	router.Mount("/graphql", routers.NewGraphQLRouter(gql))

	if err := http.ListenAndServe(":6645", router); err != nil {
		panic(err)
	}
}
