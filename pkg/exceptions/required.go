package exceptions

import "fmt"

type ErrFieldRequired struct {
	FieldName string
}

func (e ErrFieldRequired) Error() string {
	if e.FieldName != "" {
		return fmt.Sprintf("Field %s is required", e.FieldName)
	}
	return "Field is required"
}
