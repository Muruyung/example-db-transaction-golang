package exceptions_test

import (
	"coba/pkg/exceptions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrFieldRequired(t *testing.T) {
	err := exceptions.ErrFieldRequired{"Name"}
	assert.Equal(t, "Field Name is required", err.Error())
}

func TestErrFieldRequiredWithEmptyField(t *testing.T) {
	err := exceptions.ErrFieldRequired{}
	assert.Equal(t, "Field is required", err.Error())
}
