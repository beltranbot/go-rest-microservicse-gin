package httperrors

import "net/http"

// HTTPError generic http error struct
type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

// NewBadRequestError creates instance of HTTPError
func NewBadRequestError(message string) *HTTPError {
	return &HTTPError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError creates instance of HTTPError
func NewNotFoundError(message string) *HTTPError {
	return &HTTPError{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}
