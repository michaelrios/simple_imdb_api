package imdb

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

func GetYearQuery(year string) (yearQuery *bson.M, err error) {
	if year != "" {
		yearRange := strings.Split(year, ",")
		yearRangeLength := len(yearRange)

		if yearRangeLength == 1 || yearRangeLength == 2 {
			startYear, err := strconv.Atoi(yearRange[0])
			if err != nil {
				return yearQuery, errors.New(fmt.Sprintf("year:%s is invalid", yearRange[0]))
			}

			endYear := startYear
			if yearRangeLength == 2 {
				endYear, err = strconv.Atoi(yearRange[1])
				if err != nil {
					return yearQuery, errors.New(fmt.Sprintf("year:%s is invalid", yearRange[0]))
				}
			}

			yearQuery = &bson.M{"$gte": startYear, "$lte": endYear}
		} else {
			return yearQuery, errors.New("year should be formatted like 2018 or 2016,2018")
		}
	}

	return yearQuery, err
}

func GetGenreQuery(genre string) (genreQuery *string) {
	if genre != "" {
		genreQuery = &genre
	}
	return genreQuery
}

func GetQuery(yearQuery *bson.M, genreQuery *string) *bson.M {
	query := &bson.M{}

	if yearQuery != nil && genreQuery != nil {
		query = &bson.M{"year": yearQuery, "genre": genreQuery}
	} else if yearQuery != nil {
		query = &bson.M{"year": yearQuery}
	} else if genreQuery != nil {
		query = &bson.M{"genre": genreQuery}
	}

	return query
}
