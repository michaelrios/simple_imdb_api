package imdb

import (
	"github.com/michaelrios/simple_imdb_api/core"
	"net/http"
)

type API struct {
	*core.Dependencies
}

func (api *API) GetMovies(writer http.ResponseWriter, request *http.Request) {
	logger := api.GetLogger()
	logger.Info("Starting to get movies")

	// do stuff

	logger.Info("Got the movies")

	writer.Write([]byte("got movies"))
}

func (api *API) AddMovies(writer http.ResponseWriter, request *http.Request) {
	logger := api.GetLogger()
	logger.Info("Starting to add movies")

	// do stuff

	logger.Info("Added the movies")

	writer.Write([]byte("added movies"))

}
