# simple_imdb_api
## Instructions on running using docker-compose (not for production, but easy for a small project like this)
- Clone this repo (following Go conventions into "$GOPATH/src/github.com/michaelrios/simple_imdb_api")
  - or just run `go get github.com/michaelrios/simple_imdb_api`
- run `docker-compose up`
- go to `http://localhost:8080/` and you should see `Welcome to another IMDB API`
- Make a POST request to `http://localhost:8080/movies` to add the csv to be parsed and added to the DB
- Make a GET request to `http://localhost:8080/movies` to query the movies added
