package errors

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// List of errors
var (
	ErrRecordNotFound = errors.New("record not found")
)

const (
	//TxtErrHappennedTemplate ...
	TxtErrHappennedTemplate = "Có lỗi xảy ra (%s)"
)

// NOTE
// sentry-go 0.3.1 only support github.com/pkg/errors
// but in future github.com/pkg/errors will be replace by go 1.13 errors
// package, so we expect sentry-go will support built-in errors package soon

// Wrap error
func Wrap(err error, format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	zap.S().Warn(message, err)
	return errors.Wrap(err, message)
}

// New create a new error
func New(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	zap.S().Warn(message)
	return errors.New(message)
}

// NewInternalWithCode ...
func NewInternalWithCode(code string) error {
	return fmt.Errorf(TxtErrHappennedTemplate, code)
}

// Cause returns the underlying cause of the error, if possible.
func Cause(err error) error {
	return errors.Cause(err)
}

// IsRecordNotFoundError ...
func IsRecordNotFoundError(err error) bool {
	if err == ErrRecordNotFound {
		return true
	}
	if cause := errors.Cause(err); cause == ErrRecordNotFound {
		return true
	}

	return false
}

// IsDuplicatedError function check if err is DuplicatedError
func IsDuplicatedError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "pq: duplicate key") || strings.Contains(err.Error(), "pq: conflicting key")
}

// IsConstraintError function check if err is constraint
func IsConstraintError(err error) bool {
	return strings.Contains(err.Error(), "violates check constraint")
}

// ErrUnexpectedError ...
type ErrUnexpectedError struct {
	msg string
}

// NewErrUnexpectedError ...
func NewErrUnexpectedError(msg string) *ErrUnexpectedError {
	if msg == "" {
		msg = "unexpected error"
	}
	return &ErrUnexpectedError{msg}
}

// Error ...
func (e *ErrUnexpectedError) Error() string {
	return e.msg
}

// InvalidInputError ...
type InvalidInputError struct {
	Field   string
	Message string
}

// NewInvalidInputError ...
func NewInvalidInputError(field, msg string) *InvalidInputError {
	return &InvalidInputError{field, msg}
}

// Error ...
func (e *InvalidInputError) Error() string {
	return e.Message
}
