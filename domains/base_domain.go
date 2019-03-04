package domains

import (
	"errors"
	"fmt"
	"strings"
)

type BaseDomain struct {
	errors []map[string]string `json:"-"`
}

func (b *BaseDomain) IsValid() bool {
	return len(b.errors) == 0
}

func (b *BaseDomain) GetErrorMessages() []map[string]string {
	return b.errors
}

func (b *BaseDomain) GetError() error {
	messages := make([]string, 0)
	for _, errorMessage := range b.errors {
		for k, v := range errorMessage {
			messages = append(messages, fmt.Sprintf("%s: %s", k, v))
		}
	}
	return errors.New(strings.Join(messages, ","))
}

func (b *BaseDomain) addError(field string, message string) {
	b.errors = append(b.errors, map[string]string{field: message})
}
