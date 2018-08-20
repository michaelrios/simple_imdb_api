package imdb

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type Movie struct {
	ID          bson.ObjectId `json:"-" bson:"_id"`
	Rank        int           `json:"-" bson:"rank"`
	Title       string        `json:"title" bson:"title"`
	Genre       List          `json:"genre" bson:"genre"`
	Description string        `json:"description" bson:"description"`
	Director    string        `json:"-" bson:"director"`
	Actors      List          `json:"-" bson:"actors"`
	Year        int           `json:"year" bson:"year"`
	Runtime     int           `json:"runtime" bson:"runtime" csv:"Runtime (Minutes)"`
	Rating      float64       `json:"rating" bson:"rating"`
	Votes       int           `json:"-" bson:"votes"`
	Revenue     float64       `json:"-" bson:"revenue" csv:"Revenue (Millions)"`
	Metascore   int           `json:"-" bson:"metascore"`
	CreatedAt   time.Time     `json:"-" bson:"createdAt"`
}

type List []string

func (list *List) UnmarshalCSV(csv string) (err error) {
	*list = strings.Split(csv, ",")

	return nil
}
