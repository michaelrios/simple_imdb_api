package imdb_test

import (
	"github.com/michaelrios/simple_imdb_api/core"
	"github.com/michaelrios/simple_imdb_api/imdb"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"testing"
)

func TestAPI_GetMovies(t *testing.T) {
	writer := &MockResponseWriter{}
	request := &http.Request{
		URL: &url.URL{
			RawQuery: "key=val&key2=val",
		},
	}

	api := imdb.API{
		Dependencies: &core.Dependencies{
			Logger: &logrus.Logger{},
			DB: &core.MockMongo{
				Mockable: core.MockableMongo{
					Datalayer: &core.MockDatalayer{
						MockableCollection: &core.MockableCollection{},
					},
				},
			},
		},
	}

	api.GetMovies(writer, request)
}

type MockResponseWriter struct {
	Assertable AssertableMockResponseWriter
	Mockable   AssertableMockResponseWriter
}

func (w *MockResponseWriter) Header() http.Header {
	return w.Assertable.Header
}

func (w *MockResponseWriter) Write(bytes []byte) (int, error) {
	w.Assertable.Bytes = bytes
	return len(bytes), nil
}

func (w *MockResponseWriter) WriteHeader(statusCode int) {
	w.Assertable.StatusCode = statusCode
}

type AssertableMockResponseWriter struct {
	Bytes      []byte
	StatusCode int
	Header     http.Header
}
