package exceptions_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"coba/pkg/exceptions"
)

func TestCustomError(t *testing.T) {
	t.Run("Test_ERRBUSSINESS", func(t *testing.T) {
		var (
			statusCode = http.StatusBadRequest
			status     = exceptions.GetCustomErrorCodeByHttpStatusCode(statusCode)
		)

		assert.Equal(t, status, exceptions.ERRBUSSINESS)
	})

	t.Run("Test_ERRAUTHORIZED", func(t *testing.T) {
		var (
			statusCode = http.StatusUnauthorized
			status     = exceptions.GetCustomErrorCodeByHttpStatusCode(statusCode)
		)

		assert.Equal(t, status, exceptions.ERRAUTHORIZED)
	})

	t.Run("Test_ERRNOTFOUND", func(t *testing.T) {
		var (
			statusCode = http.StatusNotFound
			status     = exceptions.GetCustomErrorCodeByHttpStatusCode(statusCode)
		)

		assert.Equal(t, status, exceptions.ERRNOTFOUND)
	})

	t.Run("Test_ERRSYSTEM", func(t *testing.T) {
		var (
			statusCode = http.StatusInternalServerError
			status     = exceptions.GetCustomErrorCodeByHttpStatusCode(statusCode)
		)

		assert.Equal(t, status, exceptions.ERRSYSTEM)
	})

	t.Run("Test_ERRREPOSITORY", func(t *testing.T) {
		var (
			statusCode = http.StatusGatewayTimeout
			status     = exceptions.GetCustomErrorCodeByHttpStatusCode(statusCode)
		)

		assert.Equal(t, status, exceptions.ERRREPOSITORY)
	})

	t.Run("Test_ERRSYSTEM_Default", func(t *testing.T) {
		var (
			statusCode = http.StatusConflict
			status     = exceptions.GetCustomErrorCodeByHttpStatusCode(statusCode)
		)

		assert.Equal(t, status, exceptions.ERRSYSTEM)
	})
}
