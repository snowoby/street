package account

import (
	"net/http"
	"street/errors"
)

var (
	DuplicateEmailError = errors.HTTPError{Code: http.StatusConflict, Message: "email is already taken"}
	RecordNotMatchError = errors.HTTPError{Code: http.StatusUnauthorized, Message: "records not matched"}
	TokenNotExistsError = errors.HTTPError{Code: http.StatusUnauthorized, Message: "token not exist"}
	TokenExpiredError   = errors.HTTPError{Code: http.StatusUnauthorized, Message: "token expired"}
	WeakPasswordError   = errors.HTTPError{Code: http.StatusBadRequest, Message: "password was weak"}
	PasswordHashError   = errors.HTTPError{Code: http.StatusBadRequest, Message: "password hash error"}
)
