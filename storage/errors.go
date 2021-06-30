package storage

// ErrorType defines a type of error on file operation
type ErrorType string

// Enumeration of type of errors
var (
	ErrUnknown    ErrorType = "unknown"
	ErrNotExist   ErrorType = "not-exist"
	ErrNotDir     ErrorType = "not-dir"
	ErrPermission ErrorType = "permission"
	ErrExist      ErrorType = "exist"
	ErrNotEmpty   ErrorType = "not-empty"
	ErrTimeout    ErrorType = "timeout"
)

type Error interface {
	error
	Type() ErrorType
}
