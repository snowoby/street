package errs

import "net/http"

var (
	NotBelongsToOperator = HTTPError{Code: http.StatusForbidden, Message: "cannot operate"}
)
