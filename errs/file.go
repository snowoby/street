package errs

import "net/http"

var (
	LengthNotEqual = HTTPError{Code: http.StatusBadRequest, Message: "length and content length not "}
)
