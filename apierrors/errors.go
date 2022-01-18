package apierrors

import (
	"fmt"

	"github.com/pkg/errors"
)

// Error messages for the books-api
var (
	ErrInvalidReview        = errors.New("invalid review")
	ErrEmptyReviewMessage   = errors.New("empty review provided. Please enter a message")
	ErrEmptyReviewUser      = errors.New("empty forenames/surname provided. Please enter a valid user")
	ErrLongReviewMessage    = errors.New("review message is too long")
	ErrEmptyRequestBody     = errors.New("empty request body")
	ErrEmptyBookID          = errors.New("empty book ID in request")
	ErrEmptyReviewID        = errors.New("empty review ID in request")
	ErrUnableToReadMessage  = errors.New("failed to read request body")
	ErrUnableToParseJSON    = errors.New("failed to parse json body")
	ErrRequiredFieldMissing = errors.New("invalid book. Missing required field")
	ErrInternalServer       = errors.New("internal server error")

	APIVersionErrorMessage = "API version invalid, requires update to version in path"
)

func ErrAPIVersion(validVersions []string) error {
	return fmt.Errorf("%v. Valid versions are: %v", APIVersionErrorMessage, validVersions)
}
