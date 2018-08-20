package imdb

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type Parser struct {
	UpsertableMovies chan *Movie
	MovieRepository  *MongoRepository
	Logger           *logrus.Entry
	WG               *sync.WaitGroup
}

func (p *Parser) FillDB(numberOfWorkers int) {
	wg := sync.WaitGroup{}

	for i := 0; i < numberOfWorkers; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			for movie := range p.UpsertableMovies {
				p.UpsertMovie(movie)
			}
		}()
	}
}

func (p *Parser) UpsertMovie(movie *Movie) {
	collection := p.MovieRepository.getCollection()
	existingMovie := &Movie{}

	p.Logger.Infof("starting: %s", movie.Title)

	err := collection.
		Find(bson.M{"title": movie.Title, "year": movie.Year}).
		One(existingMovie)

	if err != nil && existingMovie == nil {
		p.Logger.WithError(err).Errorf("failed query for: %s", movie.Title)
	} else {
		if err == nil {
			err := p.MovieRepository.Update(existingMovie, movie)
			if err != nil {
				p.Logger.WithError(err).Errorf("failed updating %s", movie.Title)
			} else {
				p.Logger.Infof("updated %s", movie.Title)
			}
		} else {
			err := p.MovieRepository.Insert(movie)
			if err != nil {
				p.Logger.WithError(err).Errorf("failed inserting %s", movie.Title)
			} else {
				p.Logger.Infof("inserted %s", movie.Title)
			}
		}
	}

}
