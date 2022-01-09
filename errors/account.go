package errors

import "net/http"

var (
	DuplicateEmailError = HTTPError{Code: http.StatusConflict, Message: "email is already taken"}
	RecordNotMatchError = HTTPError{Code: http.StatusUnauthorized, Message: "records not matched"}
	TokenNotExistsError = HTTPError{Code: http.StatusUnauthorized, Message: "token not exist"}
	TokenExpiredError   = HTTPError{Code: http.StatusUnauthorized, Message: "token expired"}
	WeakPasswordError   = HTTPError{Code: http.StatusBadRequest, Message: "password was weak"}
	PasswordHashError   = HTTPError{Code: http.StatusBadRequest, Message: "password hash error"}
)
