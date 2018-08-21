package imdb

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/michaelrios/simple_imdb_api/core"
	"net/http"
	"sort"
	"sync"
)

type API struct {
	*core.Dependencies
}

func (api *API) GetMovies(writer http.ResponseWriter, request *http.Request) {
	logger := api.GetRequestLogger()
	logger.Info("Starting to get movies")

	year := request.URL.Query().Get("year")
	genre := request.URL.Query().Get("genre")

	yearQuery, err := GetYearQuery(year)
	if err != nil {
		respond(writer, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	genreQuery := GetGenreQuery(genre)

	query := GetQuery(yearQuery, genreQuery)
	logger.Info(query)

	var movies []Movie
	err = GetMovieRepository(api.DB.Copy()).
		getCollection().Find(query).Limit(10).All(&movies)
	if err != nil {
		logger.WithError(err).Error("Failed querying for movies")
		respond(writer, http.StatusBadRequest, []byte("something failed while querying"))
		return
	}

	if len(movies) == 0 {
		logger.Info("movies not found")
		respond(writer, http.StatusNotFound, []byte("movies not found"))
		return
	}

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].Rank < movies[j].Rank
	})

	movieBytes, err := json.Marshal(movies)
	if err != nil {
		logger.WithError(err).Error("Failed marshaling movies")
		respond(writer, http.StatusInternalServerError, []byte("failed marshaling movies"))
		return
	}

	respond(writer, http.StatusOK, movieBytes)
}

func (api *API) AddMovies(writer http.ResponseWriter, request *http.Request) {
	logger := api.GetRequestLogger()
	logger.Info("Starting to add movies")

	var movies []*Movie

	defer request.Body.Close()
	if err := gocsv.Unmarshal(request.Body, &movies); err != nil {
		logger.WithError(err).Error("failed to parse csv")
		respond(writer, http.StatusBadRequest, []byte("failed to parse the csv"))
		return
	}

	numberOfWorkers := 10

	parser := Parser{
		UpsertableMovies: make(chan *Movie, numberOfWorkers),
		MovieRepository:  GetMovieRepository(api.DB.Copy()),
		Logger:           logger,
		WG:               &sync.WaitGroup{},
	}
	parser.FillDB(numberOfWorkers)

	totalMovies := 0
	for _, movie := range movies {
		if movie.Title != "" {
			parser.UpsertableMovies <- movie
			totalMovies++
			logger.Infof("queueing %s", movie.Title)
		}
	}

	close(parser.UpsertableMovies)
	parser.WG.Wait()

	respond(writer, http.StatusOK, []byte(fmt.Sprintf("%d movies parsed from csv", totalMovies)))
}

func respond(writer http.ResponseWriter, status int, body []byte) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(body)
}
