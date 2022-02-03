package errs

import (
	"net/http"
)

var (
	ProfileIdentityError = HTTPError{Code: http.StatusForbidden, Message: "account and profile not matched"}
	CallDuplicateError   = HTTPError{Code: http.StatusConflict, Message: "call exists"}
)
