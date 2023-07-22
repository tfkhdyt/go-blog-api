package validator

type Error struct {
	Err error
}

func NewValidationError(err error) *Error {
	return &Error{err}
}

func (e *Error) Error() string {
	return e.Err.Error()
}
