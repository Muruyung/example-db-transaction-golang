package exceptions

import "fmt"

type ErrExists string

func (e ErrExists) Error() string {
	return fmt.Sprintf("%s already exist", string(e))
}
