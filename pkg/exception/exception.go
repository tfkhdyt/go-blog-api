package exception

type HttpError struct {
	message    string
	statusCode uint
}

func (h *HttpError) Error() string {
	return h.message
}

func (h *HttpError) StatusCode() uint {
	return h.statusCode
}

func NewHTTPError(statusCode uint, message string) *HttpError {
	return &HttpError{message, statusCode}
}
