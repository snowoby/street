package errs

import "net/http"

var (
	NotBelongsToOperator = HTTPError{code: http.StatusForbidden, message: "cannot operate"}
)
