package exception

type httpError struct {
	message    string
	statusCode uint
}

func (h *httpError) Error() string {
	return h.message
}

func NewHTTPError(statusCode uint, message string) *httpError {
	return &httpError{message, statusCode}
}
