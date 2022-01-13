package errs

import (
	"net/http"
)

var (
	ProfileIdentityError   = HTTPError{code: http.StatusForbidden, message: "account and profile not matched"}
	CallSignDuplicateError = HTTPError{code: http.StatusConflict, message: "call sign exists"}
)
