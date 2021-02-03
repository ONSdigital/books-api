package api

import (
	"context"
	"errors"
	"github.com/cadmiumcat/books-api/interfaces/datastoretest"
	"github.com/cadmiumcat/books-api/models"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const host = "http://localhost:80"

func TestEndpoints(t *testing.T) {
	t.Parallel()

	Convey("Given a POST request to add a book", t, func() {
		mockDataStore := &datastoretest.DataStoreMock{
			AddBookFunc: func(book *models.Book) {},
		}
		ctx := context.Background()

		api := Setup(ctx, host, mux.NewRouter(), mockDataStore)

		Convey("When there is no request body", func() {
			response := httptest.NewRecorder()

			body := strings.NewReader(``)
			request, err := http.NewRequest(http.MethodPost, "/books", body)
			So(err, ShouldBeNil)

			api.router.ServeHTTP(response, request)
			Convey("Then the HTTP response code is 400", func() {
				So(response.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("When the body does not contain a valid book", func() {
			response := httptest.NewRecorder()

			body := strings.NewReader(`{}`)
			request, err := http.NewRequest(http.MethodPost, "/books", body)
			So(err, ShouldBeNil)

			api.router.ServeHTTP(response, request)

			Convey("Then the HTTP response code is 400", func() {
				So(response.Code, ShouldEqual, http.StatusBadRequest)
			})
			Convey("And there AddBook function is not called", func() {
				So(len(mockDataStore.AddBookCalls()), ShouldEqual, 0)
			})
		})

		Convey("When the body contains a valid book", func() {
			response := httptest.NewRecorder()

			body := strings.NewReader(`{"title":"Girl, Woman, Other", "author":"Bernardine Evaristo" }`)
			request, err := http.NewRequest(http.MethodPost, "/books", body)
			So(err, ShouldBeNil)

			api.router.ServeHTTP(response, request)

			Convey("Then the HTTP response code is 201", func() {
				So(response.Code, ShouldEqual, http.StatusCreated)
			})
			Convey("And the AddBook function is called once", func() {
				So(len(mockDataStore.AddBookCalls()), ShouldEqual, 1)
			})
		})
	})

	Convey("Given an existing book with book id=1", t, func() {
		id := "1"
		mockDataStore := &datastoretest.DataStoreMock{
			GetBookFunc: func(id string) (*models.Book, error) {
				return &models.Book{ID: "1"}, nil
			},
		}

		ctx := context.Background()

		api := Setup(ctx, host, mux.NewRouter(), mockDataStore)
		Convey("When I send an HTTP GET request to /books/1", func() {
			response := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/books/"+id, nil)
			So(err, ShouldBeNil)

			api.router.ServeHTTP(response, request)

			Convey("Then the HTTP response code is 200", func() {
				So(response.Code, ShouldEqual, http.StatusOK)
			})
			Convey("And the GetBook function is called once", func() {
				So(len(mockDataStore.GetBookCalls()), ShouldEqual, 1)
			})
		})

	})

	Convey("Given a book that does not exist with book id=3", t, func() {
		mockDataStore := &datastoretest.DataStoreMock{
			GetBookFunc: func(id string) (*models.Book, error) {
				return nil, errors.New("error message")
			},
		}

		ctx := context.Background()

		api := Setup(ctx, host, mux.NewRouter(), mockDataStore)

		id := "3"
		Convey("When I send an HTTP GET request to /books/3", func() {
			response := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/books/"+id, nil)
			So(err, ShouldBeNil)

			api.router.ServeHTTP(response, request)
			Convey("then the HTTP response code is 404", func() {
				So(response.Code, ShouldEqual, http.StatusNotFound)
			})
			Convey("And the GetBook function is called once", func() {
				So(len(mockDataStore.GetBookCalls()), ShouldEqual, 1)
			})
		})
	})

	Convey("Given ", t, func() {
		mockDataStore := &datastoretest.DataStoreMock{
			GetBooksFunc: func() (models.Books, error) {
				return models.Books{}, nil
			},
		}

		ctx := context.Background()

		api := Setup(ctx, host, mux.NewRouter(), mockDataStore)
		Convey("When I send an HTTP GET request to /books", func() {
			response := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/books", nil)
			So(err, ShouldBeNil)

			api.router.ServeHTTP(response, request)

			Convey("then the HTTP response code is 200", func() {
				So(response.Code, ShouldEqual, http.StatusOK)
			})
			Convey("And the GetBooks function is called once", func() {
				So(len(mockDataStore.GetBooksCalls()), ShouldEqual, 1)
			})
		})
	})
}
