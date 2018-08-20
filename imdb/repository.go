package imdb

import (
	"github.com/michaelrios/simple_imdb_api/core"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoRepository struct {
	collectionName string
	dbName         string
	mongo          core.DB
}

func (orderRepository *MongoRepository) Insert(movie *Movie) error {
	movie.ID = bson.NewObjectId()
	movie.CreatedAt = time.Now()

	err := orderRepository.getCollection().
		Insert(movie)

	return err
}

func (orderRepository *MongoRepository) Update(selector *Movie, movie *Movie) error {
	err := orderRepository.getCollection().
		Update(bson.M{"_id": selector.ID}, movie)

	return err
}

func (orderRepository *MongoRepository) GetAll() (movies []Movie, err error) {
	err = orderRepository.getCollection().
		Find(nil).All(&movies)

	return movies, err
}

func GetMovieRepository(mongoSession core.DB) *MongoRepository {
	return &MongoRepository{
		collectionName: "movies",
		dbName:         "simple_imdb_api",
		mongo:          mongoSession,
	}
}

func (orderRepository *MongoRepository) getCollection() core.Collection {
	return orderRepository.mongo.
		DB(orderRepository.dbName).
		C(orderRepository.collectionName)
}
