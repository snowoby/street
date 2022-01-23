package errs

import "net/http"

var (
	LengthNotEqual = HTTPError{code: http.StatusBadRequest, message: "length and content length not "}
)
