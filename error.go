package wordsapi

import "fmt"

type ServerResponseError struct {
	Code int
}

func (e *ServerResponseError) Error() string {
	msg := ""
	switch e.Code {
	case 400:
		msg = "Bad request"
	case 401:
		msg = "Unauthorized"
	case 403:
		msg = "Forbidden"
	case 404:
		msg = "Resource not found"
	case 502:
		msg = "Bad gateway"
	case 503:
		msg = "Service unavailable"
	default:
		msg = "Unknown error"
	}
	return fmt.Sprint("Error", e.Code, ":", msg, ".")
}
