package errors

import "fmt"

type MyzaError struct {
	err           error
	message       string
	originalError error
	context       map[string]string
}

func (e MyzaError) Error() string {
	if e.originalError != nil {
		return fmt.Sprintf(`%s (%s)`, e.message, e.originalError)
	}

	return fmt.Sprintf(`%s`, e.message)
}

func (e MyzaError) Is(target error) bool {
	return e.err == target
}
