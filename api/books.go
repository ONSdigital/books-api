package api

import (
	"net/http"

	"github.com/ONSdigital/books-api/apierrors"
	"github.com/ONSdigital/books-api/models"
	"github.com/ONSdigital/books-api/pagination"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

func (api *API) addBookHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	vars := mux.Vars(request)
	apiVersion := vars["version"]
	logData := log.Data{"requested_api_version": apiVersion}

	if apiVersion == "" {
		apiVersion = api.latestVersion
	}

	if request.ContentLength == 0 {
		handleError(ctx, writer, apierrors.ErrEmptyRequestBody, nil)
		return
	}

	book := models.NewBook()
	if err := ReadJSONBody(ctx, request.Body, book); err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	logData["book"] = book

	err := book.Validate()
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	api.dataStore.AddBook(ctx, book)

	if err := WriteJSONBody(book, writer, http.StatusCreated, apiVersion); err != nil {
		handleError(ctx, writer, err, logData)
		return
	}
}

func (api *API) getBooksHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	vars := mux.Vars(request)
	apiVersion := vars["version"]
	logData := log.Data{"requested_api_version": apiVersion}

	if apiVersion == "" {
		apiVersion = api.latestVersion
	}

	offset, limit, err := api.paginator.GetPaginationValues(request)
	logData["offset"] = offset
	logData["limit"] = limit
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	books, totalCount, err := api.dataStore.GetBooks(ctx, offset, limit)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	response := models.BooksResponse{
		Items: books,
		Page: pagination.Page{
			Count:      len(books),
			Offset:     offset,
			Limit:      limit,
			TotalCount: totalCount,
		},
	}

	log.Event(ctx, "version is", log.INFO, log.Data{"version": apiVersion})

	// Update all links to contain the api version in front of the path
	for i := range response.Items {
		if response.Items[i].Links != nil {
			response.Items[i].Links.Reservations = "/" + apiVersion + response.Items[i].Links.Reservations
			response.Items[i].Links.Reviews = "/" + apiVersion + response.Items[i].Links.Reviews
			response.Items[i].Links.Self = "/" + apiVersion + response.Items[i].Links.Self
		}
	}

	if err := WriteJSONBody(response, writer, http.StatusOK, apiVersion); err != nil {
		handleError(ctx, writer, err, nil)
		return
	}

	log.Event(ctx, "successfully retrieved list of books", log.INFO)
}

func (api *API) getBookHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	vars := mux.Vars(request)
	apiVersion := vars["version"]
	id := vars["id"]
	logData := log.Data{"requested_api_version": apiVersion, "book_id": id}

	if apiVersion == "" {
		apiVersion = api.latestVersion
	}

	if id == "" {
		handleError(ctx, writer, apierrors.ErrEmptyBookID, logData)
		return
	}

	book, err := api.dataStore.GetBook(ctx, id)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	// Update links to contain the api version in front of the path
	if book != nil && book.Links != nil {
		book.Links.Reservations = "/" + apiVersion + book.Links.Reservations
		book.Links.Reviews = "/" + apiVersion + book.Links.Reviews
		book.Links.Self = "/" + apiVersion + book.Links.Self
	}

	if err := WriteJSONBody(book, writer, http.StatusOK, apiVersion); err != nil {
		handleError(ctx, writer, err, logData)
		return
	}
	log.Event(ctx, "successfully retrieved book", log.INFO, logData)
}
