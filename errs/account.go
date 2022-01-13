package errs

import "net/http"

var (
	DuplicateEmailError = HTTPError{code: http.StatusConflict, message: "email is already taken"}
	RecordNotMatchError = HTTPError{code: http.StatusUnauthorized, message: "records not matched"}
	UnauthorizedError   = HTTPError{code: http.StatusUnauthorized, message: "not login"}
	TokenNotExistsError = HTTPError{code: http.StatusUnauthorized, message: "token not exist"}
	TokenExpiredError   = HTTPError{code: http.StatusUnauthorized, message: "token expired"}
	WeakPasswordError   = HTTPError{code: http.StatusBadRequest, message: "password was weak"}
	PasswordHashError   = HTTPError{code: http.StatusBadRequest, message: "password hash error"}
)
