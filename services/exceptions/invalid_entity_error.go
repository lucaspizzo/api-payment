package exceptions

import "fmt"

type InvalidEntityError struct {
	Reason string
}

func (i *InvalidEntityError) Error() string {
	return fmt.Sprintf(i.Reason)
}
