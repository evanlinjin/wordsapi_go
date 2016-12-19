package wordsapi

import "fmt"

// Error structure used when wordsapi servers return error.
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

// Error structure used when search parameters are invalid.
type SearchParametersError struct {
	FrequencyMaxOkay bool
	FrequencyMinOkay bool
	PageOkay         bool
}

func (e *SearchParametersError) Error() (outStr string) {
	outStr = "The following parameters are invalid: "
	switch {
	case !e.FrequencyMaxOkay:
		outStr += "FrequencyMax "
		fallthrough
	case !e.FrequencyMinOkay:
		outStr += "FrequencyMin "
		fallthrough
	case !e.PageOkay:
		outStr += "Page "
	}
	return
}

type ParameterError struct {
	pType     string
	pCrrValue interface{}
	pErrValue interface{}
	pValidMin interface{}
	pValidMax interface{}
}

func (e *ParameterError) Error() string {
	switch e.pType {
	case "HasDetails":
		return fmt.Sprint(
			"Parameter", e.pType,
			"needs to be of values: [", e.pValidMax, "] .",
			"You put:", e.pErrValue, ".",
			"Value changed to:", e.pCrrValue, ".",
		)
	case "Letters":
		return fmt.Sprint(
			"Parameters", e.pType,
			"needs",
		)
	default:
		return fmt.Sprint(
			"Parameter", e.pType,
			"needs to be in range: [", e.pValidMin, ":", e.pValidMax, "] .",
			"You put:", e.pErrValue, ".",
			"Value left as:", e.pCrrValue, ".",
		)
	}

}
