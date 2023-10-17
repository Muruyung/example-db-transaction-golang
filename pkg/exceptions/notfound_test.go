package exceptions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrEntityNotFound_Error(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Entity with domain name",
			fields: fields{Name: "Domain"},
			want:   "Domain not found",
		},
		{
			name:   "Entity with empty domain name",
			fields: fields{},
			want:   "Entity not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrEntityNotFound{
				Name: tt.fields.Name,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}
