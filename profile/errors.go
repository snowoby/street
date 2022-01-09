package profile

import (
	"net/http"
	"street/errors"
)

var (
	ProfileIdentityError = errors.HTTPError{Code: http.StatusForbidden, Message: "account and profile not matched"}
)
