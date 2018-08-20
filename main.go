package main

import (
	"github.com/gorilla/mux"
	"github.com/michaelrios/simple_imdb_api/core"
	"github.com/michaelrios/simple_imdb_api/imdb"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	config := core.MustSetConfigs()

	deps := mustSetupDependencies(config)
	deps.Logger.Info("starting imdb_api")

	router := setupBaseRouter()

	addImdbToRouter(router, deps)

	deps.Logger.Info("serving imdb_api")
	server := http.Server{
		Addr:    config.ServerConfig.Addr,
		Handler: router,
	}
	defer server.Close()

	deps.Logger.Fatal(server.ListenAndServe())
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome to another IMDB API"))
}

func addImdbToRouter(router *mux.Router, deps *core.Dependencies) {
	imdbAPI := &imdb.API{Dependencies: deps}
	router.HandleFunc("/movies", imdbAPI.GetMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies", imdbAPI.AddMovies).Methods(http.MethodPost)
}

func setupBaseRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	return router
}

func mustSetupDependencies(config *core.Config) *core.Dependencies {
	deps := &core.Dependencies{}

	var err error

	deps.Logger = getLogger(&config.LoggingConfig)
	deps.DB, err = config.MongoConfig.CreateSession()
	if err != nil {
		deps.Logger.WithError(err).Fatal("failed connecting to mongo")
	}

	return deps
}

func getLogger(loggingConfig *core.LoggingConfig) *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stdout,
		Level:     logrus.Level(loggingConfig.LoggingLevel),
		Formatter: &logrus.JSONFormatter{},
	}
}
