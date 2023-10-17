package exceptions

import "fmt"

type ErrEntityNotFound struct {
	Name string
}

func (e ErrEntityNotFound) Error() string {
	if e.Name != "" {
		return fmt.Sprintf("%s not found", e.Name)
	}
	return "Entity not found"
}
