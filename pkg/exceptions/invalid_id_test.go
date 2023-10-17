package exceptions_test

import (
	"coba/pkg/exceptions"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldBeReturnInvalidErr(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Invalid id with domain name",
			fields: fields{Name: "area"},
			want:   "invalid area id",
		},
		{
			name:   "Invalid id with empty domain name",
			fields: fields{},
			want:   "invalid id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := exceptions.ErrInvalidID{
				Name: tt.fields.Name,
			}
			assert.Equalf(t, tt.want, e.Error(), "Error()")
		})
	}
}
