package errs

import "net/http"

var (
	LengthNotEqual    = HTTPError{Code: http.StatusBadRequest, Message: "length and content length not "}
	NoParts           = HTTPError{Code: http.StatusBadRequest, Message: "no parts"}
	PartsExceeded     = HTTPError{Code: http.StatusBadRequest, Message: "parts exceeded"}
	FileUploadedError = HTTPError{Code: http.StatusConflict, Message: "file is already uploaded"}
)
