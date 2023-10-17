package exceptions

import "fmt"

type ErrInvalidID struct {
	Name string
}

func (e ErrInvalidID) Error() string {
	if e.Name != "" {
		return fmt.Sprintf("invalid %s id", e.Name)
	}
	return "invalid id"
}
