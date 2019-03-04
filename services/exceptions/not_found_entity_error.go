package exceptions

import "fmt"

type NotFoundEntityError struct {
	Reason string
}

func (i *NotFoundEntityError) Error() string {
	return fmt.Sprintf(i.Reason)
}