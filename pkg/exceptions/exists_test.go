package exceptions_test

import (
	"coba/pkg/exceptions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrNameExists_Error(t *testing.T) {
	tests := []struct {
		name string
		e    exceptions.ErrExists
		want string
	}{
		{
			name: "Assertion Equal error",
			e:    exceptions.ErrExists("Floor 1"),
			want: "Floor 1 already exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.e.Error(), "Error()")
		})
	}
}
