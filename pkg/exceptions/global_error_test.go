package exceptions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"coba/pkg/exceptions"
)

func TestGlobalError(t *testing.T) {
	errMsg := "test error"
	t.Run("new_domain_error", func(t *testing.T) {
		var (
			domain = exceptions.NewDomainError(errMsg)
		)

		assert.Equal(t, errMsg, domain.Error())
	})

	t.Run("new_error_not", func(t *testing.T) {
		var (
			err = exceptions.NewErrorNot(errMsg)
		)

		assert.EqualError(t, err, errMsg)
	})
}
