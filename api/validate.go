package api

import (
	"github.com/ONSdigital/books-api/apierrors"
)

var (
	v1     = "v1"
	v2     = "v2"
	latest = v2
)

// versions provides list of versions that the API accepts
var versions = map[string]bool{
	v1: true,
	v2: true,
}

// validVersions used to provide a list of valid versions in API error response
var validVersions = []string{
	v1,
	v2,
}

func validateAPIVersion(apiVersion string) (string, error) {
	if apiVersion == "" {
		return latest, nil
	}

	if ok := versions[apiVersion]; !ok {
		return apiVersion, apierrors.ErrAPIVersion(validVersions)
	}

	return apiVersion, nil
}
