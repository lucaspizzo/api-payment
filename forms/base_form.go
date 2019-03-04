package forms

type BaseForm struct {
	errors []map[string]string `json:"-"`
}

type Operation string

const (
	Create Operation = "create"
	Update Operation = "update"
)

func (b *BaseForm) IsValid() bool {
	return len(b.errors) == 0
}

func (b *BaseForm) GetErrors() []map[string]string {
	return b.errors
}

func (b *BaseForm) addError(field string, message string) {
	b.errors = append(b.errors, map[string]string{field: message})
}