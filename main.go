package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"nino.sh/api/graphql"
	"nino.sh/api/managers"
	"nino.sh/api/metrics"
	"nino.sh/api/redis"
	"nino.sh/api/routers"
	"nino.sh/api/utils"
	"os"
)

var version = "0.2.0"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	err := godotenv.Load(".env"); if err != nil {
		panic(err)
	}

	utils.ValidateEnv()
	logrus.Infof("Running v%s of Nino API", version)
}

func main() {
	node := os.Getenv("REGION"); if node != "" {
		logrus.Infof("Running on region %s. :3", node)
	}

	logrus.WithField("bootstrap", "Postgres").Info("Connecting to PostgreSQL...")
	postgres := managers.NewPostgresManager()
	if err := postgres.GetConnection(); err != nil {
		logrus.WithField("bootstrap", "Postgres").Fatalf("Unable to connect to Postgres: %v", err)
		os.Exit(1)
	}

	logrus.WithField("bootstrap", "Redis").Info("Connecting to Redis...")
	r := redis.NewRedisClient()
	if err := r.Connect(); err != nil {
		logrus.WithField("bootstrap", "Redis").Fatalf("Unable to connect to Redis: %v", err)
		os.Exit(1)
	}

	logrus.WithField("bootstrap", "Redis").Info("Connected to Redis! :3")
	logrus.WithField("bootstrap", "GraphQL").Info("Creating GraphQL server...")
	gql := graphql.NewGraphQLManager(postgres, r)
	if err := gql.GenerateSchema(); err != nil {
		panic(err)
	}

	logrus.WithField("bootstrap", "Metrics").Info("Registering Metrics handler...")
	m := metrics.NewMetrics()
	m.Register()

	router := chi.NewRouter()
	router.Mount("/", routers.NewMainRouter())
	router.Mount("/metrics", routers.NewMetricsRouter(m))
	router.Mount("/health", routers.NewHealthRouter())
	router.Mount("/graphql", routers.NewGraphQLRouter(gql))

	logrus.WithField("bootstrap", "Http").Info("Listening at http://localhost:6645!")
	log.Fatal(http.ListenAndServe(":6645", router))
}
