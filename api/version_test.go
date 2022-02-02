package api

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/ONSdigital/books-api/apierrors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetValidVersionFunc(t *testing.T) {
	t.Parallel()
	testCtx := context.Background()

	Convey("Given a valid API version is set by service", t, func() {
		v := getValidVersion()

		Convey("When calling the getLatestVersion func", func() {
			Convey("Then the version is returned with no errors", func() {
				version, err := GetLatestVersion(testCtx, v)
				So(version, ShouldEqual, v)
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given the API version is not set", t, func() {
		v := ""

		Convey("When calling the getLatestVersion func", func() {
			Convey("Then the latest version is returned with no errors", func() {
				version, err := GetLatestVersion(testCtx, v)
				So(version, ShouldEqual, defaultLatest)
				So(err, ShouldResemble, apierrors.ErrAPIVersion(ValidVersions))
			})
		})
	})

	Convey("Given the API version is not valid", t, func() {
		v := getInvalidVersion()

		Convey("When calling the getLatestVersion func", func() {
			Convey("Then the latest version is returned with no errors", func() {
				version, err := GetLatestVersion(testCtx, v)
				So(version, ShouldEqual, defaultLatest)
				So(err, ShouldResemble, apierrors.ErrAPIVersion(ValidVersions))
			})
		})
	})
}

// getValidVersion retrieves a valid version
func getValidVersion() string {
	for version, isValid := range versions {
		if isValid {
			return version
		}
	}

	return defaultLatest
}

// getInvalidVersion retrieves an invalid version
func getInvalidVersion() (v string) {
	for v, isValid := range versions {
		if !isValid {
			return v
		}
	}

	version, err := strconv.Atoi(strings.TrimPrefix("v", v))
	if err != nil {
		// Just return a known invalid version, v0
		return "v0"
	}

	return "v" + strconv.Itoa(version+1)
}
