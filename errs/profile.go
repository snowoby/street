package errs

import (
	"net/http"
)

var (
	NoProfiles           = HTTPError{Code: http.StatusNotFound, Message: "no profiles available"}
	ProfileIdentityError = HTTPError{Code: http.StatusForbidden, Message: "account and profile not matched"}
	CallDuplicateError   = HTTPError{Code: http.StatusConflict, Message: "call exists"}
)
