package exceptions

import "fmt"

type DomainError struct {
	message string
}

func NewDomainError(message string) DomainError {
	return DomainError{message: message}
}

func (d DomainError) Error() string {
	return d.message
}

func NewErrorNot(message string) error {
	err := NewDomainError(message)
	return fmt.Errorf("%s", err)
}
