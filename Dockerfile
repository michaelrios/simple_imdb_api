FROM golang:1.10.1-alpine3.7
RUN apk add --update make

# define work directory
RUN mkdir -p /go/src/github.com/michaelrios/simple_imdb_api
WORKDIR /go/src/github.com/michaelrios/simple_imdb_api

# Adding source files
COPY . /go/src/github.com/michaelrios/simple_imdb_api

# Ensure code quality and build
RUN make test && go build

EXPOSE 8080
CMD ["./simple_imdb_api"]
