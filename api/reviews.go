package api

import (
	"net/http"

	"github.com/ONSdigital/books-api/apierrors"
	"github.com/ONSdigital/books-api/models"
	"github.com/ONSdigital/books-api/pagination"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

func (api *API) addReviewHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	vars := mux.Vars(request)
	apiVersion := vars["version"]
	bookID := vars["id"]
	logData := log.Data{"requested_api_version": apiVersion, "book_id": bookID}

	apiVersion, err := validateAPIVersion(apiVersion)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	if bookID == "" {
		handleError(ctx, writer, apierrors.ErrEmptyBookID, logData)
		return
	}

	if request.ContentLength == 0 {
		handleError(ctx, writer, apierrors.ErrEmptyRequestBody, logData)
		return
	}

	// Confirm that book exists. If bookID not found, then a review cannot be added!
	_, err = api.dataStore.GetBook(ctx, bookID)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	review := models.NewReview(bookID)

	if err := ReadJSONBody(ctx, request.Body, review); err != nil {
		handleError(ctx, writer, apierrors.ErrInvalidReview, logData)
		return
	}

	logData["review"] = review

	err = review.Validate()
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	api.dataStore.AddReview(ctx, review)

	if err := WriteJSONBody(review, writer, http.StatusCreated, apiVersion); err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

}

func (api *API) getReviewsHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	vars := mux.Vars(request)
	apiVersion := vars["version"]
	bookID := vars["id"]
	logData := log.Data{"requested_api_version": apiVersion, "book_id": bookID}

	apiVersion, err := validateAPIVersion(apiVersion)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	offset, limit, err := api.paginator.GetPaginationValues(request)
	logData["offset"] = offset
	logData["limit"] = limit
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	if bookID == "" {
		handleError(ctx, writer, apierrors.ErrEmptyBookID, logData)
		return
	}

	// Confirm that book exists. If bookID not found, then do not check for the reviews
	_, err = api.dataStore.GetBook(ctx, bookID)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	reviews, totalCount, err := api.dataStore.GetReviews(ctx, bookID, offset, limit)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	response := models.ReviewsResponse{
		Items: reviews,
		Page: pagination.Page{
			Count:      len(reviews),
			Offset:     offset,
			Limit:      limit,
			TotalCount: totalCount,
		},
	}

	// Update all links to contain the api version in front of the path
	for i := range response.Items {
		response.Items[i].Links.Book = "/" + apiVersion + response.Items[i].Links.Book
		response.Items[i].Links.Self = "/" + apiVersion + response.Items[i].Links.Self
	}

	if err := WriteJSONBody(response, writer, http.StatusOK, apiVersion); err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	log.Event(ctx, "successfully retrieved review", log.INFO, logData)
}

func (api *API) getReviewHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	vars := mux.Vars(request)
	apiVersion := vars["version"]
	bookID := vars["id"]
	reviewID := vars["reviewID"]
	logData := log.Data{"requested_api_version": apiVersion, "book_id": bookID, "review_id": reviewID}

	apiVersion, err := validateAPIVersion(apiVersion)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	if bookID == "" {
		handleError(ctx, writer, apierrors.ErrEmptyBookID, logData)
		return
	}

	if reviewID == "" {
		handleError(ctx, writer, apierrors.ErrEmptyReviewID, logData)
		return
	}

	// Confirm that book exists. If bookID not found, then do not check for the review
	_, err = api.dataStore.GetBook(ctx, bookID)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	review, err := api.dataStore.GetReview(ctx, reviewID)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	// Update all links to contain the api version in front of the path
	if review != nil && review.Links != nil {
		review.Links.Book = "/" + apiVersion + review.Links.Book
		review.Links.Self = "/" + apiVersion + review.Links.Self
	}

	if err := WriteJSONBody(review, writer, http.StatusOK, apiVersion); err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	log.Event(ctx, "successfully retrieved review", log.INFO, logData)
}

func (api *API) updateReviewHandler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	// TODO all endpoints, whether get, post, put, delete should handle API version! Previously only handled get
	// TODO also need to handle failure scenarios to return header with version!
	apiVersion := mux.Vars(request)["version"]

	bookID := mux.Vars(request)["id"]
	reviewID := mux.Vars(request)["reviewID"]

	logData := log.Data{"book_id": bookID, "review_id": reviewID}

	if bookID == "" {
		handleError(ctx, writer, apierrors.ErrEmptyBookID, logData)
		return
	}

	if reviewID == "" {
		handleError(ctx, writer, apierrors.ErrEmptyReviewID, logData)
		return
	}

	// Confirm that book exists. If bookID not found, or there's another error, then return
	_, err := api.dataStore.GetBook(ctx, bookID)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	// Confirm that the review exists. If reviewID not found, or there's another error, then return
	_, err = api.dataStore.GetReview(ctx, reviewID)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	review := &models.Review{User: models.User{}}
	if err := ReadJSONBody(ctx, request.Body, review); err != nil {
		handleError(ctx, writer, apierrors.ErrInvalidReview, logData)
		return
	}

	logData["review"] = review

	err = api.dataStore.UpdateReview(ctx, reviewID, review)
	if err != nil {
		handleError(ctx, writer, err, logData)
		return
	}

	writer.Header().Set("Content-Type", "application/vnd.books."+apiVersion+"+json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

}
