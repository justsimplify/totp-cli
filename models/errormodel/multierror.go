package errormodel

import (
	"fmt"
	"strings"
)

// MultiError multi error type
type MultiError struct {
	Errors []error
}

// Error return error string
func (e *MultiError) Error() string {
	var errArr []string
	for _, v := range e.Errors {
		errArr = append(errArr, v.Error())
	}
	return fmt.Sprintf("%v", strings.Join(errArr, "\n"))
}

func (e *MultiError) Append(v error) {
	e.Errors = append(e.Errors, v)
}

func (e *MultiError) CheckIfNoError() bool {
	return len(e.Errors) == 0
}
