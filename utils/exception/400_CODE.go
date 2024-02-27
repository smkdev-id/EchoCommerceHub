package exception

import "fmt"

type NotFoundError struct {
	Message string
	ID      uint
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %d", e.Message, e.ID)
}
