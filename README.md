# simple_imdb_api
## Instructions on running using docker-compose (not for production, but easy for a small project like this)
- Clone this repo (following Go conventions into "$GOPATH/src/github.com/michaelrios/simple_imdb_api")
  - or just run `go get github.com/michaelrios/simple_imdb_api`
- install dep `brew install dep`
- run `dep ensure` from the simple_imdb_api directory
- install docker https://www.docker.com/get-started
- run `docker-compose up` from the simple_imdb_api directory
- go to `http://localhost:8080/` and you should see `Welcome to another IMDB API`
- Make a POST request to `http://localhost:8080/movies` to add the csv to be parsed and added to the DB
  - ie `curl --request POST --data-binary @/Users/michaelrios/Downloads/IMDB-Movie-Data.csv http://localhost:8080/movies`
- Make a GET request to `http://localhost:8080/movies` to query the movies added
  - `curl --request GET http://localhost:8080/movies\?year\=2006\&genre\=Action`
  - `curl --request GET http://localhost:8080/movies\?year\=2006,2014\&genre\=Action`
  - `curl --request GET http://localhost:8080/movies\?genre\=Action`
  - `curl --request GET http://localhost:8080/movies\?year\=2006`
  - `curl --request GET http://localhost:8080/movies\?year\=2006,2010`
