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
	logger := getLogger()
	logger.Info("starting imdb_api")

	deps := setupDependencies(logger)

	router := setupBaseRouter()

	addImdbToRouter(router, deps)

	logger.Info("serving imdb_api")
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	defer server.Close()

	logger.Fatal(server.ListenAndServe())

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

func setupDependencies(logger *logrus.Logger) *core.Dependencies {
	deps := &core.Dependencies{}
	deps.SetBaseLogger(logger)

	// add mysql

	return deps
}

func getLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stdout,
		Level:     logrus.InfoLevel,
		Formatter: &logrus.JSONFormatter{},
	}
}
