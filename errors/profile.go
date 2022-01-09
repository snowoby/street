package errors

import (
	"net/http"
)

var (
	ProfileIdentityError   = HTTPError{Code: http.StatusForbidden, Message: "account and profile not matched"}
	CallSignDuplicateError = HTTPError{Code: http.StatusConflict, Message: "call sign exists"}
)
