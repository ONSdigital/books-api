package api

import (
	"context"

	"github.com/ONSdigital/books-api/apierrors"
)

var (
	v1            = "v1"
	v2            = "v2"
	defaultLatest = v1
)

// versions provides list of versions that the API accepts
var versions = map[string]bool{
	v1: true,
	v2: true,
}

// validVersions used to provide a list of valid versions in API error response
var ValidVersions = []string{
	v1,
	v2,
}

func GetLatestVersion(ctx context.Context, version string) (string, error) {
	if ok := versions[version]; !ok {
		return defaultLatest, apierrors.ErrAPIVersion(ValidVersions)
	}

	return version, nil
}
