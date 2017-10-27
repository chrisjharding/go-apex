package apex

type Error interface {
	error
	GetType() string
}

type baseError struct {
	message   string
	errorType string
}

func NewError(message, errorType string) Error {
	return &baseError{
		errorType: errorType,
	}
}

func WrapErrorType(err error, errorType string) Error {
	return &baseError{
		message:   err.Error(),
		errorType: errorType,
	}
}

func (e *baseError) GetType() string {
	return e.errorType
}

func (e *baseError) Error() string {
	return e.message
}
