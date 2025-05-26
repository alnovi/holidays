package validator

type ValidateError struct {
	Fields map[string]string
}

func NewValidateError() *ValidateError {
	return &ValidateError{Fields: make(map[string]string)}
}

func NewValidateErrorWithMessage(key, message string) *ValidateError {
	e := NewValidateError()
	e.Fields[key] = message
	return e
}

func (e *ValidateError) Error() string {
	return "Unprocessable Entity"
}
