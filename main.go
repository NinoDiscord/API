package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"nino.sh/api/graphql"
	"nino.sh/api/managers"
	"nino.sh/api/routers"
	"nino.sh/api/utils"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	err := godotenv.Load(".env"); if err != nil {
		panic(err)
	}

	utils.ValidateEnv()

	node := os.Getenv("REGION"); if node != "" {
		logrus.Infof("Running on region %s. :3", node)
	}

	logrus.Info("Now bootstrapping API...")
	postgres := managers.NewPostgresManager()
	err = postgres.GetConnection(); if err != nil {
		logrus.Fatalf("Unable to connect to Postgres: %v", err)
		os.Exit(1)
	}

	gql := graphql.NewGraphQLManager(postgres)
	if err := gql.GenerateSchema(); err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Mount("/", routers.NewMainRouter())
	router.Mount("/health", routers.NewHealthRouter())
	router.Mount("/graphql", routers.NewGraphQLRouter(gql))

	logrus.Info("Listening at http://localhost:6645!")
	log.Fatal(http.ListenAndServe(":6645", router))
}
